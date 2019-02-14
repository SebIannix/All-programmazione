import java.util.*;
public class Soluzione {
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		int x = in.nextInt();
		while (x > 0 && x < 100) {
			if (x <= 6) {
				System.out.print("fascia 3");
				break;
			} else if (x >= 7 && x <= 11) {
				System.out.print("fascia 7");
				break;
			} else if (x >= 12 && x <= 15) {
				System.out.print("fascia 12");
				break;
			} else if (x >= 16 && x <= 17) {
				System.out.print("fascia 16");
				break;
			} else if (x >= 18) {
				System.out.print("fascia 18");
				break;
			}
		}
	}
}
