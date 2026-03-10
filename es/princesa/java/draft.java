import java.util.ArrayList;
import java.util.Scanner;

public class draft {
    ArrayList<int> fila = new ArrayList<>();
    int referencial;
    
    @Override
    public String toStringDraft(){
        String str = null;
        str += "[ ";
        
        for(int i=0; i<fila.size();i++){
            str += fila.get(i);
            if(fila.get(i)==e){
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
        String ns = scan.nextLine();
        String es = scan.nextLine();
        int n = Integer.parseInt(ns);
        int e = Integer.parseInt(es);
        referencial = e;
        
        for(int i=0; i<n; i++){
            fila.add(i);
        }

        System.out.println(toStringDraft());


        for(int i=0; i<n; i++){
            if(fila.size()==1)
                break;
            
            if(fila.get(i)==e){
                if((e+1) == n-1){
                    fila.remove(e+1);
                    e = fila.get(0);
                }

                else if((e) == n-1){
                    fila.remove(0);
                    e = fila.get(1);
                }
                
                else{
                    fila.remove(e+1);
                    e = fila.get(e+2);
                }
            }
            referencial = e;

        }
    }
}

