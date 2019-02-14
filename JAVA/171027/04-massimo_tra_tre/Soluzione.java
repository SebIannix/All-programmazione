import java.util.*;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		int x = in.nextInt();
		int y = in.nextInt();
		int z = in.nextInt();
		if (x > y) {
			if (x > z) {
				System.out.println(x);
			} else
				System.out.println(z);
		} else if (y > z) {
			System.out.println(y);
		} else
			System.out.println(z); 
	}
}
