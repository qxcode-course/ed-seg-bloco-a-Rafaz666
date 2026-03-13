package es;

public class IO{

    public static void print(String str){
        System.out.println(str);
    }
    
    public static Object scan(String str){
        
        if(str.matches("-?\\d+")){
            int i = Integer.parseInt(str);
            return i;
        }

        if(str.matches("-?\\d+\\.\\d+")){
            double d = Double.parseDouble(str);
            return d;
        }

        return str;
    }
} 