import java.util.*;
public class soluzione {

	public static int spazi;

	public static void main(String[]args){
		Scanner in = new Scanner(System.in);
		
		int i = 0;
		
		
		while (in.hasNextInt()) {
			int x = in.nextInt();
			if (i % 2 == 0) {
				fila(x);
				i++;
			} else {
				colonna(x);
				i++;
			}
		}
	}
	
	public static void fila(int x) {
		if (spazi > 0) {
			for (int i = 0; i < spazi - 1; i++)
				System.out.print(" ");
			System.out.print("+");
		}
		for (int i = 0; i < x; i++)
			System.out.print("-");
		System.out.println("+");
		spazi += x;
	}
	
	public static void colonna(int x) {
		for (int i = 0; i < x; i++) {
			for (int k = 0; k < spazi; k++)
				System.out.print(" ");
			System.out.println("|");
		}
		spazi++;
	}
}
