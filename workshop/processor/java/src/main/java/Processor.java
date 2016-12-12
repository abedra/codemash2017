import redis.clients.jedis.Jedis;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.stream.Stream;

public class Processor {
    public static HashMap<String, Actor> counts  = new HashMap<>();
    private final String logFile;
    private static long threshold;
    private static Jedis jedis;

    public Processor(String logFile, long threshold) {
        this.logFile = logFile;
        Processor.threshold = threshold;
        jedis = new Jedis("localhost");
    }

    public static LogEvent getData(String line) {
        String[] parts = line.split(" ");
        return new LogEvent(parts[0], parts[5].replace("\"", ""), parts[8], parts[6]);
    }

    public static void updateCounts(LogEvent event) {
        Actor actor = counts.get(event.getAddress());

        if (actor == null) {
            actor = new Actor(event.getAddress());
        }

        if (event.getMethod().equals("POST") && event.getResponse().equals("200") && event.getUri().equals("/")) {
            actor.incrementFailedLogins();
            counts.put(event.getAddress(), actor);
        }

        if (Integer.parseInt(event.getResponse()) >= 400) {
            actor.incrementFailedResponse(event.getResponse());
            counts.put(event.getAddress(), actor);
        }
    }

    public static void check(String address, Actor actor) {
        if (actor.getFailedLogins() >= threshold) {
            System.out.println("Blacklisting " + actor.getName() + ". Failed Login Threshold 10, Actual: " + actor.getFailedLogins());
            jedis.set(actor.getName() + ":repsheet:ip:blacklisted", "failedLogin");
        }

        if (actor.getTotalFailedResponses() >= threshold) {
            System.out.println("Blacklisting " + actor.getName() + ". Failed Response Threshold 10, Actual: " + actor.getTotalFailedResponses());
            jedis.set(actor.getName() + ":repsheet:ip:blacklisted", "failedResponse");
        }
    }

    public void process() throws IOException {
        Path path = Paths.get(logFile);

        try (Stream<String> lines = Files.lines(path)) {
            lines.map(Processor::getData).forEach(Processor::updateCounts);
        }

        counts.forEach(Processor::check);
    }

    public void cleanup() {
        jedis.close();
    }

    public static void main(String[] args) throws IOException {
	Processor processor = new Processor(args[0], Long.parseLong(args[1]));
	processor.process();
	processor.cleanup();
    }
}
