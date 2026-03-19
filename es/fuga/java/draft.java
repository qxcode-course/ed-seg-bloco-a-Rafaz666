import java.util.*;

public class draft {
    static Scanner scan = new Scanner(System.in);
    static boolean terminou;
    
    public static void main(String args[]) {
        int h = scan.nextInt();
        int p = scan.nextInt();
        int f = scan.nextInt();
        int d = scan.nextInt();

        if(d == 1){
            for(f=f; terminou == false; f++){
                if(f > 15)
                    f=0;
                
                if(f == h){
                    System.out.println("S");
                    terminou = true;
                }
                else if(f == p){
                    System.out.println("N");
                    terminou = true;
                }
            }
        }

        if(d == -1){
            for(f=f; terminou == false; f--){
                if(f < 0)
                    f=15;
                
                if(f == h){
                    System.out.println("S");
                    terminou = true;
                }
                else if(f == p){
                    System.out.println("N");
                    terminou = true;
                }
            }
        }
    }

}

