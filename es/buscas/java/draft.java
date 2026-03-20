import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class draft {
    static Scanner scan =  new Scanner(System.in);
    static ArrayList<String> consul = new ArrayList<>();
    static ArrayList<String> input = new ArrayList<>();
    
    public static void main(String args[]) {
        int qtdInput = scan.nextInt();
        scan.nextLine();
        input = IntStream.range(0, qtdInput).mapToObj(i -> scan.next()).collect(Collectors.toCollection(ArrayList::new));
        int qtdConsul = scan.nextInt();
        scan.nextLine();
        consul = IntStream.range(0, qtdConsul).mapToObj(i -> scan.next()).collect(Collectors.toCollection(ArrayList::new));
        String str = "";
        
        for(int i = 0; i<input.size(); i++){
            
            if(consul.contains(input.get(i))){
                cont++;
                str += cont + " ";
            }
        }

        System.out.println(str);

        

    }
}
