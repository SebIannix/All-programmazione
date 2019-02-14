import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		int x = in.nextInt();
		
		
		for (int i = 1; i <= x; i++) {
		
		
		if (i % 3 == 0 && i % 7 == 0)
			System.out.println("Tico Taco");
		else if (i % 3 == 0)
			System.out.println("Tico");
		else if (i % 7 == 0)
			System.out.println("Taco");
			
		if (i % 3 != 0 && i % 7 != 0)
			System.out.println(i);
		}
		
	}
}
