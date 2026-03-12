import java.util.Scanner;

public class draft {
    static Scanner scan = new Scanner (System.in);
    static Integer referencial;
    static Integer[] fila;

    public static void org(){
        for(Integer i=0; i<fila.length; i++){
            if(fila[i] == null){
                fila[i] = fila[i+1];
                fila[i+1] = null;
            }
        }
    }

    public static void select(){
        for(Integer i=0; i<fila.length; i++){
            if(referencial.equals(fila[i])){
                if(referencial == fila.length){
                    referencial = fila[i+1];
                    kill(fila[0]);
                }
                else if (referencial == fila.length-1){
                    referencial = fila[0];
                    kill(fila[i+1]);
                }
                else{
                    referencial = fila[i+2];
                    kill(fila[i+1]);
                }
            }
        }
    }

    public static void kill(Integer index){
        fila[index] = null;
        org();
    }

    public static String toStringDraft(){
        String str = "";
        str += "[ ";
        for(Integer i=0; i<fila.length; i++){
            str += fila[i] + " ";
        }
        str += "]";

        return str;
    }

    public static void main(String[] args){
        Integer qtd = scan.nextInt();
        referencial = scan.nextInt();
        fila = new Integer[qtd];
        
        for(Integer i=0; i<fila.length; i++){
            fila[i] = i+1;
        }

        while(fila.length!=1){
            select();
            System.out.println(toStringDraft());
        }


        

    }




}

