import java.util.*;
public class soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		int x = in.nextInt();
		String op = in.next();
		int y = in.nextInt();
		if (op.equals("+")) {
			System.out.println(x + y);
		} if (op.equals("-")) {
			System.out.println(x - y);
		} if (op.equals("*")) {
			System.out.println(x * y);
		} if (op.equals("/")) {
			if (y == 0)
				System.out.println("errore");
			else if (x%y != 0)
				System.out.println(x/y + " con resto " + x%y);
			else
				System.out.println(x/y);
		}
	}
}
