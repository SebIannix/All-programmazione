import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		double x = in.nextDouble();
		double y = in.nextDouble();
		double z = in.nextDouble();
		double add = (x + y + z);
		int somma = (int) (x + y + z);
		double resto = (add - somma)*100;
		int rf = (int)resto;
		
		System.out.println(somma + " euro e " + Math.abs(rf) + " centesimi");
		
	}
}
