import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    public static void main(String args[]) {
        int qtd = scan.nextInt();
        int[] fila = IntStream.range(0, qtd).map(i -> scan.nextInt()).toArray();
        int qtdquit = scan.nextInt();
        int filaquit[] = IntStream.range(0, qtdquit).map(i -> scan.nextInt()).toArray();




        String str = "saida: \n";
        str += qtd + "\n";
        str += Arrays.stream(fila).mapToObj(x -> String.valueOf(x)).collect(Collectors.joining(" ")).toString().trim() + "\n";
        str += qtdquit + "\n";
        str += Arrays.stream(filaquit).mapToObj(x -> String.valueOf(x)).collect(Collectors.joining(" ")).toString().trim();
        System.out.println(str);

    }
}
