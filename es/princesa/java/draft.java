import java.util.ArrayList;
import java.util.Scanner;

public class draft {
    static int[] fila;
    static int referencial;
    
    public static String toStringDraft(){
        String str = "";
        str += "[ ";
        
        for(int i=0; i<fila.size();i++){
            str += fila[i];
            if(fila[i]==referencial){
                str += ">";
            }
            
            if(i!=fila.size()-1){
                str += " ";
            }
            
            else{
                str += " ]";
            }
        }

        return str;
    }
    public static void main(String args[]) {
        Scanner scan = new Scanner(System.in);
        String ns[] = scan.nextLine().split(" ");
        int n = Integer.parseInt(ns[0]);
        int e = Integer.parseInt(ns[1]);
        referencial = e;
        
        for(int i=0; i<n; i++){
            fila[i] = i+1;
        }

        System.out.println(fila.toString());

        while(fila.size()!=1){
        
            if(fila.get(e+1) == n-1){
                fila.remove(e+1);
                e = fila.get(0);
            }

            else if(fila.get(e) == n-1){
                fila.remove(0);
                e = fila.get(0);
            }
                
            else{
                fila.remove(e+1);
                e = fila.get(e+1);
            }
            
            referencial = e;
            System.out.println(fila.toString());
        }
    }
}

