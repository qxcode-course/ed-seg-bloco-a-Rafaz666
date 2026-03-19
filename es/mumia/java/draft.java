import java.util.Scanner;
public class draft {
    
    public static void print(String str){
        System.out.println(str);
    }
    public static void main(String args[]) {
        Scanner scan = new Scanner (System.in);
        String str[] = (scan.nextLine() + " " + scan.nextLine()).split(" ");
        int age = Integer.parseInt(str[1]);

        if(age<12){
            print(str[0] + " eh crianca");
        }
        else if(age<18){
            print(str[0] + " eh jovem");
        }
        else if(age<65){
            print(str[0] + " eh adulto");
        }
        else if(age<1000){
            print(str[0] + " eh idoso");
        }
        else{
            print(str[0] + " eh mumia");
        }
    }
}

