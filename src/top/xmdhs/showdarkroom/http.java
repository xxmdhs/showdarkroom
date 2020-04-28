package top.xmdhs.showdarkroom;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.nio.charset.StandardCharsets;

public class http {
    private final URL url;
    public http(URL url){
        this.url = url;
    }
    public String getjson() {
        try {
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("GET");
        connection.setUseCaches(false);
        connection.setConnectTimeout(5000);
        connection.setReadTimeout(5000);
        connection.setRequestProperty("Accept", "*/*");
        connection.setRequestProperty("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36");
      try(BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream(), StandardCharsets.UTF_8))) {
        String current;
        StringBuilder json = new StringBuilder();
        while ((current = in.readLine()) != null) {
            json.append(current);
        }
        return json.toString();
      }
        } catch (IOException e) {
            e.printStackTrace();
            return "";
        }

    }
    public String getcid(String json){
       try {
           int a = json.indexOf("\"cid\":\"");
           int b = json.indexOf("\"},\"data\"");
           return json.substring(a+7,b);
       }catch (Exception e){
           e.printStackTrace();
           return "";
       }

    }
    public static int getStartCid() {
        try {
        URL url = new URL("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom");
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("GET");
        connection.setUseCaches(false);
        connection.setConnectTimeout(5000);
        connection.setReadTimeout(5000);
        connection.setRequestProperty("Accept", "*/*");
        connection.setRequestProperty("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36");
        BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream(), StandardCharsets.UTF_8));
        String current;
        while ((current = in.readLine()) != null) {
            if(current.contains("id=\"darkroommore\" cid=\"")){
                int a = current.indexOf("id=\"darkroommore\" cid=\"");
                int b = current.indexOf("\">更多</a></div></span>");
                return Integer.parseInt(current.substring(a+23,b));
            }
        }
        return 10000;
        } catch (IOException e) {
            e.printStackTrace();
            return 0;
        }
    }
}
