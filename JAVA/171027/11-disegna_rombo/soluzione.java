import java.util.*;
public class soluzione {
	public static void main(String[]args) {
		Scanner in = new Scanner(System.in);
		
		int n = in.nextInt();
		
		for (int i=1; i<n; i++) {
			for (int j=1; j<=2*n-1; j++)
				if ((j<=n-1) || (j>=n+1))
					System.out.print(' ');
				else
					System.out.print('*');
			System.out.println();
		}
		
	}
}
