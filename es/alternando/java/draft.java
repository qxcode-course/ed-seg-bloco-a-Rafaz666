import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer[] fila;
    static Integer qtd;
    static Integer e;
    static Integer f;
    static boolean end = false;

    public static Integer limit(){
        return fila.length;
    }

    public static void org(){
        fila = Arrays.stream(fila).filter(x -> x != null).toArray(Integer[]::new);
    }

    public static int edgeR(int i){
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


    public static int edgeL(int i){
        if(0>i-1 && 0>i-2){
            return 0;
        }
        else if(0>i-1){
            return 1;
        }
        else if( 0<i-1 && 0>i-2){
            return 2;
        }
        return -1;
    }


    public static void killRight(){
        for(int i=0; i<fila.length; i++){

            if(fila.length<=1){
                end = true;
                break;
            }
            
            else if(fila[i].equals(e)){
                switch (edgeR(i)) {
                    case 0 -> {
                        fila[i+1] = null;
                        e = fila[i+2];
                        org();
                        return;
                    }
                    case 1 -> {
                        fila[0] = null;
                        e = fila[1];
                        org();
                        return;
                    }
                    case 2 -> {
                        fila[i+1] = null;
                        e = fila[0];
                        org();
                        return;
                    }
                    default -> System.out.println("deu pau");
                }
            }
        }
    }

    public static void killLeft(){
        for(int i = fila.length-1; i>=0; i--){
            if(fila.length<=1){
                end = true;
                break;
            }
            
            else if(fila[i].equals(e)){
                switch (edgeL(i)) {
                    case 0 -> {
                        fila[i-1] = null;
                        e = fila[i-2];
                        org();
                        return;
                    }
                    case 1 -> {
                        fila[fila.length-1] = null;
                        e = fila[fila.length-2];
                        org();
                        return;
                    }
                    case 2 -> {
                        fila[i-1] = null;
                        e = fila[fila.length-1];
                        org();
                        return;
                    }
                    default -> System.out.println("deu pau");
                }
            }
        }
    }

    public static void draftString(){
        String str = "[ ";

        if(f.equals(-1)){
            for (Integer fila1 : fila) {
                if (fila1.equals(e)) {
                    str += "<" + fila1 + " ";
                } else {
                    str += fila1 + " ";
                }
            }
        }
        else{
            for (Integer fila1 : fila) {
                if (fila1.equals(e)) {
                    str += fila1 + "> ";
                } else {
                    str += fila1 + " ";
                }
            }
        }

        str += "]";

        System.out.println(str);
    }
    public static void main(String[] args){
        qtd = scan.nextInt();
        e = scan.nextInt();
        f = scan.nextInt();
        
        fila  = IntStream.rangeClosed(1,qtd).boxed().toArray(Integer[]::new);
        
        if(f.equals(-1)){
            for( int i=0; i<fila.length; i++){
                if(fila[i]%2!=0){
                    fila[i] = -1*i;
                }
            }
        }
        else{
            for( int i=0; i<fila.length; i++){
                if(fila[i]%2==0){
                    fila[i] = -1*i;
                }
            }
        }

        while(end==false){
            draftString();
            
            if(f.equals(-1))
                killLeft();

            else
                killRight();
        }
    }
}

