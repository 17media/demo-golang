public class AccountHelper {
    public static List<Account> getAccounts() {
        return [SELECT Id, Name FROM Account];
    }
}