using System;
using System.IO;
using System.Collections.Concurrent;
using StackExchange.Redis;
using Microsoft.Extensions.CommandLineUtils;

namespace ConsoleApplication
{
    public class Processor
    {
        public string logfile { get; set; }
        public int threshold { get; set; }
        public ConcurrentDictionary<string, int> entries;
        public Processor(string logfile, int threshold)
        {
            this.logfile = logfile;
            this.threshold = threshold;
            this.entries = new ConcurrentDictionary<string, int>();
        }

        public string[] ReadFile()
        {
            return File.ReadAllLines(this.logfile);
        }

        public void Run()
        {
            foreach (var line in ReadFile())
            {   
                var parts  = line.Split(' ');
                var address = parts[0];
                var method = parts[5].Replace("\"", "");
                var uri = parts[6];                
                var responseCode = parts[8];

                if (method == "POST" && responseCode == "200")
                {
                    this.entries.AddOrUpdate(address, 1, (k,v) => v + 1);
                }
            }

            if (this.entries.Count > 0)
            {
                var connetion = ConnectionMultiplexer.Connect("127.0.0.1");
                IDatabase redis = connetion.GetDatabase();
                foreach (var entry in this.entries)
                {
                    Console.WriteLine("Blacklisting {0}. Threshold {1}, Actual {2}", entry.Key, this.threshold, entry.Value);
                    var key = entry.Key + ":repsheet:ip:blacklisted";
                    redis.StringSet(key, "Failed login processor");
                }
            }
        }
        public static void Main(string[] args)
        {
            CommandLineApplication cli = new CommandLineApplication(throwOnUnexpectedArg: false);
            CommandOption logFile = cli.Option("-f | --file <filename>", "The path and name of the logfile", CommandOptionType.SingleValue);
            CommandOption threshold = cli.Option(" -t | --threshold", "The blacklist threshold", CommandOptionType.SingleValue);
            cli.OnExecute(() => 
            {
                if (logFile.HasValue() && threshold.HasValue())
                {
                    var processor = new Processor(logFile.Value(), Int32.Parse(threshold.Value()));
                    processor.Run();
                }
                else
                {
                    Console.WriteLine("Must supply a logfile and threshold");
                }

                return 0;
            });
            cli.Execute(args);
        }
    }
}
