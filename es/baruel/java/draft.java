import java.util.*;
import java.util.stream.IntStream;
public class draft{
    static Scanner scan = new Scanner(System.in);
    static int album;
    static int qtd;
    static int[] firstInput;

    static void outPut(){
        if(falta().length==0){
            System.out.println("N");
        }
        else{
            int faltam[] = falta();
            for(int i=0; i<faltam.length; i++){
                if(i<faltam.length-1)
                    System.out.print(faltam[i] + " ");
                    else
                        System.out.println(faltam[i]);
            }
        }

        if(repet()==0){
            System.out.println("N");
        }
        else{
            System.out.println(repet());
        }

    }

    static int[] falta(){
        int[] falta = IntStream.rangeClosed(0, album-1).filter(x -> Arrays.stream(firstInput).noneMatch(y -> y == x)).toArray();
        return falta;
    }
    
    static int repet(){
        int repet = (int) Arrays.stream(firstInput).distinct().map(x -> (int) Arrays.stream(firstInput).filter(y -> y==x ).count()-1).sum();
        return repet;
    }



    public static void main(String[] args){
        album = scan.nextInt();
        qtd = scan.nextInt();
        scan.nextLine();
        firstInput = Arrays.stream(scan.nextLine().split(" ")).mapToInt(s -> Integer.parseInt(s)).limit(qtd).toArray();

        outPut();
    }
}