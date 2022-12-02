var Twitter = function() {
  this.tweets = [];
  this.follows = {};
};
Twitter.prototype.postTweet = function(userId, tweetId) {
  this.tweets.push({ userId: userId, tweetId: tweetId });
};
Twitter.prototype.getNewsFeed = function(userId) {
  var result = [];
  var follows = this.follows[userId] || [];
  follows.push(userId);
  for (var i = this.tweets.length - 1; i >= 0; i--) {
    if (follows.indexOf(this.tweets[i].userId) !== -1) {
      result.push(this.tweets[i].tweetId);
      if (result.length === 10) {
        break;
      }
    }
  }
  return result;
};
Twitter.prototype.follow = function(followerId, followeeId) {
  if (!this.follows[followerId]) {
    this.follows[followerId] = [];
  }
  this.follows[followerId].push(followeeId);
};
Twitter.prototype.unfollow = function(followerId, followeeId) {
  if (!this.follows[followerId]) {
    return;
  }
  var index = this.follows[followerId].indexOf(followeeId);
  if (index !== -1) {
    this.follows[followerId].splice(index, 1);
  }
};