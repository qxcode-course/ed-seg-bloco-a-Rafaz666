import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer[] fila;
    static Integer qtd;
    static Integer e;
    static int limite = fila.length;

    public static void org(){
        fila = Arrays.stream(fila).filter(x -> x != null).toArray(Integer[]::new);
    }

    public static int edge(int i){
        if(fila[i+1]<limite && fila[i+2]<limite){
            return 0;
        }
        else if(fila[i+1]>=limite){
            return 1;
        }
        else if(fila[i+1]<limite && fila[i+2]>=limite){
            return 2;
        }
        return -1;
    }

    
    public static void main(String[] args){
        qtd = scan.nextInt();
        e = scan.nextInt();
        fila  = IntStream.rangeClosed(1,qtd).boxed().toArray(Integer[]::new);

        for(Integer i : fila)
            System.out.print(i + " ");

        for(int i=0; i<fila.length; i++){
            if(fila[i].equals(e)){
                if(edge(i)==0){
                    Sy
                }
            }
        }
        
        org();

        for(Integer i : fila)
            System.out.print(i + " ");

    }
}
