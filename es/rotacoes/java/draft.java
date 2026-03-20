import java.util.*;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    
    public static void main(String args[]) {
        int inputqtd = scan.nextInt();
        int qtd = scan.nextInt();
        int[] input = IntStream.rangeClosed(1, inputqtd).toArray();

        for(int i=0; i<qtd; i++){
            int j = input.length-1;
            int valor = input[input.length-1];
            
            while(j>0){
               input[j] = input[j-1];
               j--;
            }

            input[0] = valor;
        }

        
        String str = "[ ";
        for(int i : input){
            str += i + " ";
        }
        str += "]";

        System.out.println(str);
    }
}
