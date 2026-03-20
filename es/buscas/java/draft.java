import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    static ArrayList<String> consul = new ArrayList<>();
    static ArrayList<String> input = new ArrayList<>();

    public static int countCountain(String str){
        int cont = 0;
        for(int i=0; i<input.size(); i++){
            if(str.equals(input.get(i)))
                cont++;
        }

        return cont;
    }
    
    public static void main(String args[]) {
        int qtdInput = scan.nextInt();
        scan.nextLine();
        input = IntStream.range(0, qtdInput).mapToObj(i -> scan.next()).collect(Collectors.toCollection(ArrayList::new));
        int qtdConsul = scan.nextInt();
        scan.nextLine();
        consul = IntStream.range(0, qtdConsul).mapToObj(i -> scan.next()).collect(Collectors.toCollection(ArrayList::new));
        String str = "";
        for(String s : consul){
            str += countCountain(s) + " ";
        }
        str = str.trim();

        System.out.println(str);
    

        

    }
}
