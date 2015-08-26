import "auth/SingleAccountAuth.sol";

// The 'admin' field is set through the constructor.
// There is a function to check if the caller is 'admin', and also a way for the
// current admin to pass the admin role over to another account.
/// @title AdminAuth
/// @author Andreas Olofsson (andreas@erisindustries.com)
contract AdminAuth is SingleAccountAuth {

    function AdminAuth(address _admin){
        admin = _admin;
    }

}