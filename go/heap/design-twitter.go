package main

import (
	"container/heap"
)
type tweetHeap []Tweet

func (h tweetHeap) Len() int{
	return len(h)
}

func (h tweetHeap) Less(i int, j int) bool {
	return h[i].time > h[j].time
}

func (h tweetHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *tweetHeap) Push(x any) {
	*h = append(*h, x.(Tweet))
}
func (h *tweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
type Tweet struct {
	ID int
	userID int
	time int
}
type User struct {
	ID int
	ownTweets []Tweet
	follows map[int]*User
	
}
type Twitter struct {
   users map[int]*User 
   counter int
}


func TwitterConstructor() Twitter {
   return Twitter{users: make(map[int]*User), counter: 0} 
}

func NewUser(id int) *User {
	return &User{
		ID: id,
		follows: make(map[int]*User),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
   tweet := Tweet{ID : tweetId, userID: userId, time: this.counter}
   this.counter++
   usr, ok := this.users[userId] 
   if !ok {
	   usr = NewUser(userId)
	   this.users[userId] = usr
   }
   usr.ownTweets = append(usr.ownTweets, tweet)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
   usr , ok := this.users[userId] 
   if !ok {
	return []int{}
   }
   allTweets := usr.ownTweets
   for _, f := range usr.follows{
		allTweets = append(allTweets, f.ownTweets...)
   }
   h := tweetHeap(allTweets) 
   heap.Init(&h)
   res := []int{}
   for range 10{
		if h.Len() == 0{
			break
		}
		t := heap.Pop(&h).(Tweet)
		res = append(res, t.ID)
   }

   return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
   follower, ok := this.users[followerId] 
   if !ok {
	follower = NewUser(followerId)
	this.users[followerId] = follower
   }
   followee , ok:= this.users[followeeId]
   if !ok {
	followee = NewUser(followeeId)
	this.users[followeeId] = followee
   }
   follower.follows[followeeId] = followee
   
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
   follower := this.users[followerId] 
   _ , ok := follower.follows[followeeId]
   if !ok {
	return
   }
   delete(follower.follows, followeeId)
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */