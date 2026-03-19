import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer[] fila;
    static Integer qtd;
    static Integer e;
    static boolean end = false;

    public static Integer limit(){
        return fila.length;
    }

    public static void org(){
        fila = Arrays.stream(fila).filter(x -> x != null).toArray(Integer[]::new);
    }

    public static int edge(int i){
        if(i+1<limit() && i+2<limit()){
            return 0;
        }
        else if(i+1>=limit()){
            return 1;
        }
        else if(i+1<limit() && i+2>=limit()){
            return 2;
        }
        return -1;
    }

    public static void kill(){
        for(int i=0; i<fila.length; i++){

            if(fila.length<=1){
                end = true;
                break;
            }
            
            else if(fila[i].equals(e)){
                switch (edge(i)) {
                    case 0 -> {
                        fila[i+1] = null;
                        e = fila[i+2];
                        org();
                        break;
                    }
                    case 1 -> {
                        fila[0] = null;
                        e = fila[1];
                        org();
                        break;
                    }
                    case 2 -> {
                        fila[i+1] = null;
                        e = fila[0];
                        org();
                        break;
                    }
                    default -> System.out.println("deu pau");
                }
            }
        }
    }

    
    public static void draftString(){
        String str = "[ ";
        
        for(int i=0; i<fila.length; i++){
            if(fila[i].equals(e))
                str += fila[i] + "> ";
            else
                str += fila[i] + " ";
        }

        str += "]";

        System.out.println(str);
    }
    public static void main(String[] args){
        qtd = scan.nextInt();
        e = scan.nextInt();
        fila  = IntStream.rangeClosed(1,qtd).boxed().toArray(Integer[]::new);

        while(end==false){
            draftString();
            kill();
        }
    }
}
