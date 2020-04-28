package top.xmdhs.showdarkroom;

import java.io.*;
import java.net.MalformedURLException;
import java.net.URL;
import java.nio.charset.StandardCharsets;

public class a {
    public static void main(String[] args){
        System.out.println("开始爬取");
        Thread pa1 = new Thread(new Pa(616526,411018));
        Thread pa2 = new Thread(new Pa(411018,205510));
        Thread pa3 = new Thread(new Pa(205510,0));
        pa1.start();
        pa2.start();
        pa3.start();
    }
}

class Pa implements Runnable{
    private final int STRATCID;
    private final int ENDCID;

    public Pa(int STRATCID, int ENDCID){
        this.STRATCID = STRATCID;
        this.ENDCID = ENDCID;
    }

    public void run() {
        int cid = STRATCID;
       while (cid > ENDCID){
            try {
                System.out.println(STRATCID + ": "+ cid);
                http Http = new http(new URL("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom&cid="+ cid +"&ajaxdata=json"));
                String json = Http.getjson();
                if(cid == Integer.parseInt(Http.getcid(json))){
                    cid = ENDCID;
                    continue;
                }
                cid = Integer.parseInt(Http.getcid(json));
                try(BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(new FileOutputStream(cid+".txt"), StandardCharsets.UTF_8))) {
                    writer.write(json);
                } catch (IOException e) {
                    e.printStackTrace();
                }
                Thread.sleep(500);
            } catch (MalformedURLException | InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
