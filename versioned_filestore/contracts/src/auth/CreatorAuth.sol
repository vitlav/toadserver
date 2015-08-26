import "auth/SingleAccountAuth.sol";

// The contract creator becomes admin.
// There is a function to check if the caller is 'admin', and also a way for the
// current admin to pass the admin role to another account.
/// @title CreatorAuth
/// @author Andreas Olofsson (andreas@erisindustries.com)
contract CreatorAuth is SingleAccountAuth {

    function CreatorAuth(){
        admin = msg.sender;
    }

}