import java.util.HashMap;

public class Actor {
    private long failedLogins = 0;
    private HashMap<String, Long> failedResponses = new HashMap<>();
    private final String name;

    public Actor(String name) {
        this.name = name;
    }

    public void incrementFailedLogins() {
        failedLogins += 1;
    }

    public void incrementFailedResponse(String responseCode) {
        Long previous = failedResponses.get(responseCode);
        failedResponses.put(responseCode, previous == null ? 1 : previous + 1);
    }

    public long getFailedLogins() {
        return failedLogins;
    }

    public Long getTotalFailedResponses() {
         return failedResponses.values().stream().mapToLong(i -> i).sum();
    }

    public String getName() {
        return name;
    }
}
