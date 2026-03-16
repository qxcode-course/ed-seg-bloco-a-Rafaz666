import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer[] fila;
    static Integer qtd;
    static Integer e;

    public static void org(){
        fila = Arrays.stream(fila).filter(x -> x != null).toArray(Integer[]::new);
    }

    public static void kill(){
        for(int i=0; i<fila.length;i++){
            if(fila.length>1 && e.equals(fila[i])){
                if(fila[i+1]<fila.length){
                    fila[i+1] = null;
                    org();
                }
                else{
                    fila[0] = null;
                    org();
                }
            }
            else{
                break;
            }
        }
    }
    
    public static void main(String[] args){
        qtd = scan.nextInt();
        e = scan.nextInt();
        fila  = IntStream.rangeClosed(1,qtd).boxed().toArray(Integer[]::new);

        
    }




}

