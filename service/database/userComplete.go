package database

func (db *appdbimpl) CompleteUser_info_by_ids (requestingUser uint64, requestedUser uint64, checkBanned bool) (CompleteUser, error) {
	/* get the requestedUser CompleteUser resource, as viewed by requestingUser 
	   Optionally, it can check that the searched user had not banned the searching user
	   through the parameter checkBanned */
	// if used with checkBanned == true, Implements GET /users/{userID} 
	
	var result CompleteUser

	if checkBanned {
		isbanned, err := db.isBanned_by_uid(requestingUser, requestedUser)

		if err != nil {
			return result, err
		}
		if isbanned {
			return result, &BannedRequestingUserError{}
		}
	}

	err := db.c.QueryRow(`
		Select 
			U.id, U.username,
			nof.numberOfFollowers, af.accountsFollowed,
			nop.numberOfPosts,
			CASE WHEN EXISTS B THEN 1 ELSE 0 END AS isBanned,
			CASE WHEN EXISTS F THEN 1 ELSE 0 END AS isFollowed
		FROM
			(select * from users where id = ? ) as U,
			(select count(*) as numberOfFollowers from follows where followedId = ? ) as nof,
			(select count(*) as accountsFollowed  from follows where followerId = ? ) as af,
			(select count(*) as numberOfPosts from photos where authorId = ? ) as nop,
			(select * from bans where bannedId= ? and bannerId = ? ) as B,
			(select * from follows where followedId= ? and followerId = ? ) as F;
		`, requestedUser, requestedUser, requestedUser, requestedUser, 
		   requestedUser, requestingUser, requestedUser, requestingUser).Scan(&result)
	return result, err
}

func (db *appdbimpl) getUserList (requestingUser uint64, prompt string) ([]CompleteUser, error) {
	/* returns at most 24 CompleteUsers records with the requested prompt 
	in their username among those who have not banned the requesteduser */

	// Implements GET /users/
	
	var result []CompleteUser
	redUsers, err := db.ReducedUser_usernameLike(requestingUser, prompt)
	if (err != nil) {return result,err}

	for _,redUser := range redUsers {
		complUser, err := db.CompleteUser_info_by_ids(requestingUser, redUser.id, false)
		if (err != nil) {return result,err}
		result = append(result, complUser)
	}
	return result, err
}



