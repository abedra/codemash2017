#include <fstream>
#include <string>
#include <sstream>
#include <iostream>
#include <vector>
#include <unordered_map>
#include <redis3m/redis3m.hpp>
#include "boost/program_options.hpp"

using namespace std;

int main(int argc, char **argv) {
  namespace po = boost::program_options;
  po::options_description description("Processor Usage");

  description.add_options()
    ("help,h", "Display this help message")
    ("logfile,l", po::value<std::string>()->default_value(""), "Logfile location")
    ("threshold,t", po::value<int>()->default_value(10), "Threshold");

  po::variables_map vm;
  po::store(po::command_line_parser(argc, argv).options(description).run(), vm);
  po::notify(vm);

  if (vm.count("help")) {
    cout << description;
    return 0;
  }

  if (vm["logfile"].as<string>().empty()) {
    cout << "Must supply a logfile" << endl;
    cout << description;
    return 0;
  }

  int threshold = vm["threshold"].as<int>();
  string logfile = vm["logfile"].as<string>();

  ifstream infile(logfile);
  string line;
  unordered_map<string, int> actors;

  while (getline(infile, line)) {
    istringstream iss(line);
    string part;
    vector<string> parts;

    while (iss >> part) {
      parts.push_back(part);
    }

    parts[5].erase(0,1);

    if (parts[5] == "POST" && parts[8] == "200") {
      actors[parts[0]]++;
    }
  }

  if (actors.size() > 0) {
    redis3m::connection::ptr_t conn;
    try {
      conn = redis3m::connection::create();
    } catch (redis3m::unable_to_connect) {
      cout << "Unable to connect to redis, exiting." << endl;
      return -1;
    }

    for (auto &actor : actors) {
      if (actor.second > threshold) {
        ostringstream command;
        command << actor.first << ":repsheet:ip:blacklisted";
        conn->run(redis3m::command("SET") << command.str() << "log parser");
        cout << "Blacklisting " << actor.first << ". Threshold: " << threshold << ", Actual: " << actor.second << endl;
      }
    }
  }

  return 0;
}
