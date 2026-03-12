import java.util.Arrays;
import java.util.Scanner;
public class draft {
    public static void main(String args[]) {
        Scanner scan = new Scanner(System.in);
        String[] str = (scan.nextLine() + " " + scan.nextLine() + " " + scan.nextLine()).split(" ");
        double lados[] = Arrays.stream(str).mapToDouble(s -> Double.parseDouble(s)).toArray();
        double p = (lados[0] + lados[1] + lados[2])/2;
        double area = Math.sqrt(p * (p-lados[0]) * (p-lados[1]) * (p-lados[2])); 
        
        System.out.println(String.format("%.2f", area));
    }
}


