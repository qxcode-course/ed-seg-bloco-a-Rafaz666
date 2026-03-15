import java.util.*;
import java.util.stream.IntStream;
public class draft{
    static Scanner scan = new Scanner(System.in);
    static int album;
    static int qtd;
    static int[] firstInput;

    static void outPut(){

        if(repet().length==0){
            System.out.println("N");
        }
        else{
            int[] repetem = repet();
            for(int i=0; i<repetem.length; i++){
                if(i<repetem.length-1)
                    System.out.print(repetem[i] + " ");
                    else
                        System.out.println(repetem[i]);
            }
        }

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
    }

    static int[] falta(){
        int[] figualbum = IntStream.rangeClosed(1, album).toArray();
        int[] falta = Arrays.stream(figualbum).filter( x -> Arrays.stream(firstInput).noneMatch(y -> y==x)).toArray();
        return falta;
    }
    
    static int[] repet(){
        int[] repet = Arrays.stream(firstInput).distinct().flatMap(x -> IntStream.generate(() -> x).limit(Arrays.stream(firstInput).filter(y -> y==x).count()-1)).toArray();
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