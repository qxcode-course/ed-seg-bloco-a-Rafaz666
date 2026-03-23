<<<<<<< Updated upstream
import java.util.*;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    static int[] eixoX;
    static int[] eixoY;
    static  int qtd;
    static String s;

    public static void move(String str){
        switch (s) {
            case "R" -> {
                eixoX[0] = eixoX[0] + 1;
                for(int i=1; i<qtd; i++)
                    eixoX[i] = eixoX[i-1] - 1;
            }
            case "L"-> {
                eixoX[0] = eixoX[0] - 1;
                for(int i=1; i<qtd; i++)
                    eixoX[i] = eixoX[i-1] + 1;
            }
            case "D"-> {
                eixoY[0] = eixoY[0] + 1;
                for(int i=1; i<qtd; i++)
                    eixoY[i] = eixoY[i-1] - 1;
            }
            case "U" -> {
                eixoY[0] = eixoY[0] - 1;
                for(int i=1; i<qtd; i++)
                    eixoY[i] = eixoY[i-1] + 1;
            }
            default -> System.out.println("erro");       
        }

    }
    public static void main(String args[]) {
        qtd = scan.nextInt();
         s = scan.nextLine();
        eixoX = new int[qtd];
        eixoY = new int[qtd];
        
        for(int i=0; i<qtd; i++){
            eixoX[i] = scan.nextInt();
            eixoY[i] = scan.nextInt();
        }

        move(s);

        String str = "";
        for(int i=0; i<qtd; i++)
            str += eixoX[i] + " " + eixoY[i] + "\n";
        
        System.out.println(str);
    }
}

=======
import java.util.*;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    static int[] eixoX;
    static int[] eixoY;
    static  int qtd;
    static String s;

    public static void move(String str){
        for(int i=qtd-1; i>0; i--){
            
            eixoX[i] = eixoX[i-1];
        }
    }
    public static void main(String args[]) {
        qtd = scan.nextInt();
        scan.nextLine();
        s = scan.nextLine();
        eixoX = new int[qtd];
        eixoY = new int[qtd];
        
        for(int i=0; i<qtd; i++){
            eixoX[i] = scan.nextInt();
            eixoY[i] = scan.nextInt();
        }

        move(s);
    }
}

>>>>>>> Stashed changes
