package models

// Void message uchun struktura
type Void struct{}

// Id message uchun struktura
type Id struct {
	UserID string `json:"user_id" db:"user_id"`
}

// CreateRequest message uchun struktura
type CreateRequest struct {
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
}

// UserResponse message uchun struktura
type UserResponse struct {
	Id          string `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Username    string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
	Bio         string `json:"bio" db:"bio"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

// LoginRequest message uchun struktura
type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// LoginResponse message uchun struktura
type LoginResponse struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	UserID       string `json:"user_id" db:"user_id"`
}

// GetProfileResponse message uchun struktura
type GetProfileResponse struct {
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	PhoneNumber    string `json:"phone_number" db:"phone_number"`
	Username       string `json:"username" db:"username"`
	Nationality    string `json:"nationality" db:"nationality"`
	Bio            string `json:"bio" db:"bio"`
	ProfileImage   string `json:"profile_image" db:"profile_image"`
	FollowersCount int32  `json:"followers_count" db:"followers_count"`
	FollowingCount int32  `json:"following_count" db:"following_count"`
	PostsCount     int32  `json:"posts_count" db:"posts_count"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

// UpdateProfileRequest message uchun struktura
type UpdateProfileRequest struct {
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	Username     string `json:"username" db:"username"`
	Nationality  string `json:"nationality" db:"nationality"`
	Bio          string `json:"bio" db:"bio"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	Phone        string `json:"phone" db:"phone"`
}

// Filter message uchun struktura
type Filter struct {
	Role      string `json:"role" db:"role"`
	Page      int32  `json:"page" db:"page"`
	Limit     int32  `json:"limit" db:"limit"`
	FirstName string `json:"first_name" db:"first_name"`
}

// UserResponses message uchun struktura
type UserResponses struct {
	Users []UserResponse `json:"users" db:"users"`
}

// ChangePasswordRequest message uchun struktura
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" db:"current_password"`
	NewPassword     string `json:"new_password" db:"new_password"`
}

// ChangePasswordResponse message uchun struktura
type ChangePasswordResponse struct {
	Message string `json:"message" db:"message"`
}

// URL message uchun struktura
type URL struct {
	URL string `json:"url" db:"url"`
}

// Ids message uchun struktura
type Ids struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
}

// Followings message uchun struktura
type Followings struct {
	Ids []Ids `json:"ids" db:"ids"`
}

// Follower message uchun struktura
type Follower struct {
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
}

// Followers message uchun struktura
type Followers struct {
	Followers []Follower `json:"followers" db:"followers"`
}

// DFollowRes corresponds to the DFollowRes message
type DFollowRes struct {
	FollowerID   string `json:"follower_id" db:"follower_id"`
	FollowingID  string `json:"following_id" db:"following_id"`
	UnfollowedAt string `json:"unfollowed_at" db:"unfollowed_at"`
}

// Count corresponds to the Count message
type Count struct {
	Description string `json:"description" db:"description"`
	Count       int64  `json:"count" db:"count"`
}

// FollowReq corresponds to the FollowReq message
type FollowReq struct {
	FollowingID string `json:"following_id" db:"following_id"`
}

// FollowRes corresponds to the FollowRes message
type FollowRes struct {
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
	FollowedAt  string `json:"followed_at" db:"followed_at"`
}

//----------------------------------------------------------------------

// Tweet struct corresponds to the Tweet message
type Tweet struct {
	UserID    string `json:"user_id" db:"user_id"`
	Hashtag   string `json:"hashtag" db:"hashtag"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	ImageURL  string `json:"image_url" db:"image_url"`
	CreatedAt string `json:"created_at" db:"created_at"`
	LikeCount int64  `json:"like_count" db:"like_count"`
}

// TweetResponse struct corresponds to the TweetResponse message
type TweetResponse struct {
	UserID    string `json:"user_id" db:"user_id"`
	Hashtag   string `json:"hashtag" db:"hashtag"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	ImageURL  string `json:"image_url" db:"image_url"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

// UpdateATweet struct corresponds to the UpdateATweet message
type UpdateATweet struct {
	ID      string `json:"id" db:"id"`
	Hashtag string `json:"hashtag" db:"hashtag"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

// Url struct corresponds to the Url message
type Url struct {
	TweetID string `json:"tweet_id" db:"tweet_id"`
	URL     string `json:"url" db:"url"`
}

// Message struct corresponds to the Message, message
type Message struct {
	Message string `json:"message" db:"message"`
}

// UserID struct corresponds to the UserId message
type UserID struct {
	ID string `json:"id" db:"id"`
}

type Error struct {
	Message string `json:"message" db:"message"`
}

// Tweets struct corresponds to the Tweets message
type Tweets struct {
	Tweets []TweetResponse `json:"tweets" db:"tweets"`
	Limit  int64           `json:"limit" db:"limit"`
	Offset int64           `json:"offset" db:"offset"`
}

// TweetId struct corresponds to the TweetId message
type TweetId struct {
	ID string `json:"id" db:"id"`
}

// TweetFilter struct corresponds to the TweetFilter message
type TweetFilter struct {
	Limit   int64  `json:"limit" db:"limit"`
	Offset  int64  `json:"offset" db:"offset"`
	Hashtag string `json:"hashtag" db:"hashtag"`
	Title   string `json:"title" db:"title"`
}

// User struct corresponds to the User message
type User struct {
	UserID         string `json:"user_id" db:"user_id"`
	FirstName      string `json:"first_name" db:"first_name"`
	Username       string `json:"username" db:"username"`
	Bio            string `json:"bio" db:"bio"`
	ProfileImage   string `json:"profile_image" db:"profile_image"`
	FollowersCount int32  `json:"followers_count" db:"followers_count"`
	FollowingCount int32  `json:"following_count" db:"following_count"`
	PostsCount     int32  `json:"posts_count" db:"posts_count"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

// LikeReq struct corresponds to the LikeReq message
type LikeReq struct {
	TweetID string `json:"tweet_id" db:"tweet_id"`
}

// LikeRes struct corresponds to the LikeRes message
type LikeRes struct {
	UserID  string `json:"user_id" db:"user_id"`
	TweetID string `json:"tweet_id" db:"tweet_id"`
	LikedAt string `json:"liked_at" db:"liked_at"`
}

// DLikeRes struct corresponds to the DLikeRes message
type DLikeRes struct {
	UserID    string `json:"user_id" db:"user_id"`
	TweetID   string `json:"tweet_id" db:"tweet_id"`
	UnlikedAt string `json:"unliked_at" db:"unliked_at"`
}

// TweetTitles struct corresponds to the TweetTitles message
type TweetTitles struct {
	Titles []string `json:"titles" db:"titles"`
}

// Comment struct corresponds to the Comment message
type Comment struct {
	ID        string `json:"id" db:"id"`
	TweetID   string `json:"tweet_id" db:"tweet_id"`
	Content   string `json:"content" db:"content"`
	LikeCount int64  `json:"like_count" db:"like_count"`
}

// CommentRes struct corresponds to the CommentRes message
type CommentRes struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	TweetID   string `json:"tweet_id" db:"tweet_id"`
	Content   string `json:"content" db:"content"`
	LikeCount int64  `json:"like_count" db:"like_count"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

// UpdateAComment struct corresponds to the UpdateAComment message
type UpdateAComment struct {
	ID      string `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
}

// CommentId struct corresponds to the CommentId message
type CommentId struct {
	ID string `json:"id" db:"id"`
}

// CommentFilter struct corresponds to the CommentFilter message
type CommentFilter struct {
	TweetID string `json:"tweet_id" db:"tweet_id"`
}

// Comments struct corresponds to the Comments message
type Comments struct {
	Comments []CommentRes `json:"comments" db:"comments"`
}

// CommentLikeReq struct corresponds to the CommentLikeReq message
type CommentLikeReq struct {
	CommentID string `json:"comment_id" db:"comment_id"`
}
