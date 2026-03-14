
import java.util.Arrays;
import java.util.Scanner;

public class draft {
    static Scanner scan = new Scanner(System.in);
    public static void main(String args[]) {
        int vez = scan.nextInt();
        int[] jog = new int[vez];
        int a;
        int b;
        int menor;
        int ganhador = 0;


        for(int i=0; i<vez; i++){
            a = scan.nextInt();
            b = scan.nextInt();
            if(a>=10 && b>=10){
                jog[i] = Math.abs(a-b);
            }
        }

        if (Arrays.stream(jog).allMatch(n -> n==jog[0])) {
            System.out.println("sem ganhador");
            return;
        }
        
        menor = Arrays.stream(jog).filter(n -> n!=0).min().getAsInt();

        for(int i=0; i<vez; i++){
            if(menor==jog[i]){
                ganhador = i;
            }
        }

        System.out.println(ganhador);

    }
}


