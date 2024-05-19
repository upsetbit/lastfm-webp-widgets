class User {
  constructor() {
    this.listeningStatus = document.getElementById('listening-status');
    this.stats           = document.getElementById('user-stats');
    this.url             = document.getElementById('user-url');
  }

  setUrl(username) {
    const url = 'last.fm/user/' + username;
    this.url.textContent = url;

    return url;
  }

  setStats(scrobbles, creation) {
    const stat = scrobbles + ' scrobbles since ' + creation;
    this.stats.textContent = stat;

    return stat;
  }

  setListeningStatusNowPlaying() {
    const status = 'now playing';
    this.listeningStatus.textContent = status;

    return status;
  }

  setListeningStatusLastPlayed(relativeTime) {
    const status = 'last played ' + relativeTime + ' ago';
    this.listeningStatus.textContent = status;

    return status;
  }
}

const user = new User();
