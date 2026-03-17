import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer[] fila;
    static Integer qtd;
    static Integer e;
    static boolean terminou = false;

    public static void org(){
        fila = Arrays.stream(fila).filter(x -> x != null).toArray(Integer[]::new);
    }

    public static Integer outEdge(Integer i){
    
        if(i+1<fila.length && i+2<fila.length){
            return null;
        }
        else if(i+1<fila.length && i+2 >= fila.length){
            return 2;
        }
        else if(i+1>=fila.length){
            return 1;
        }

        return null;

    }

    public static void kill(){
        for(int i=0; i<fila.length;i++){
            if(fila.length > 1){
                if(e.equals(fila[i])){
                    if(outEdge(i)==null){
                        fila[i+1] = null;
                        e = fila[i+2];
                        org();
                    }
                    else if(outEdge(i).equals(2)){
                        fila[i+1] = null;
                        e = fila[0];
                        org();
                    }
                    else{
                        fila[0] = null;
                        e = fila[1];
                        org();
                    }
                }
            }
            else{
                terminou = true;
                break;
            }

        }
    }
    
    public static void main(String[] args){
        qtd = scan.nextInt();
        e = scan.nextInt();
        fila  = IntStream.rangeClosed(1,qtd).boxed().toArray(Integer[]::new);
        
        while(terminou==false){
            String str = "[ ";
            
            for(int i=0; i<fila.length; i++){
                if(i<fila.length && e.equals(fila[i]))
                    str += fila[i] + "> ";
                else if(i<fila.length)
                    str += fila[i] + " ";
            }
            
            str += "]";
            System.out.println(str);
            kill();


        }

        
    }




}

