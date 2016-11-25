public class LogEvent {
    private final String method;
    private final String address;
    private final String response;
    private final String uri;

    public LogEvent (String address, String method, String response, String uri) {
        this.method = method;
        this.address = address;
        this.response = response;
        this.uri = uri;
    }

    public String getMethod() {
        return method;
    }

    public String getAddress() {
        return address;
    }

    public String getResponse() {
        return response;
    }

    public String getUri() {
        return uri;
    }
}
