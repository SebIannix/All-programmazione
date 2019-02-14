import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		int x = in.nextInt();
		int y = 1;
		System.out.print(x + " ");
		while (x != 1) {
			if (x % 2 == 0) {
				x = x/2;
				y++;
				System.out.print(x + " ");
			} else {
				x = (x*3 + 1);
				y++;
				System.out.print(x + " ");
			}
		}
		System.out.println(y);
	}
}
