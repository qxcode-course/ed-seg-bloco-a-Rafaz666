import java.util.*;
public class draft {
    static Scanner scan = new Scanner(System.in);
    public static void main(String args[]) {
        int casais = 0;
        int tamanho = scan.nextInt();
        scan.nextLine();
        int especies[] = new int[tamanho]; 
        especies = Arrays.stream(scan.nextLine().split(" ")).mapToInt(s -> Integer.parseInt(s)).toArray();

        for(int i=0, x=0; i<tamanho; i++){
            if(especies[i]==0)
                continue;
            
            else{
                x = especies[i];
                especies[i] = 0;
            }
            
            for(int j=0; j<tamanho; j++){
                if (especies[j] == -x){
                    casais++;
                    especies[j]=0;
                    break;
                }
            }

        }

        System.out.println(casais);
    }
}

