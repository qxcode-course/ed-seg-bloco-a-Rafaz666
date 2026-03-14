import java.util.*;
public class draft {
    static Scanner scan =  new Scanner(System.in);
    static Integer[] figu = new Integer[qtd];
    static Integer album = scan.nextInt();
    static Integer qtd = scan.nextInt();

    public static void ordenar(){
        Integer[] ord = new Integer[album];
        Integer menor;

        for(Integer i=0; i<album; i++){
            menor = Arrays.stream(figu).min(Integer::compareTo).get();
            ord[i] = menor;
            
            for(Integer j=0; j<qtd; j++){
                if(menor.equals(figu[j])){
                    figu[j] = null;
                }
            }
        }

    }
    public static void main(String args[]) {
        
        for(Integer i=0;i<qtd; i++){
            figu[i] = scan.nextInt();
        }



    
    }
}

