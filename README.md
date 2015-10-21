# Github Digest

Aggregate open pull requests, report on closed pull requests and user activity.


### Usage Info
```
NAME:
   github-digest - Report on github pull request activity

USAGE:
   ./github-digest [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cutoff "21"        Days of pulls to consider
   --closed-cutoff "1"  Days of merged pulls to consider
   --oauth              Github OAuth token [$GITHUB_OAUTH_TOKEN]
   --json               Dump JSON instead of HTML
   --help, -h           show help
   --version, -v        print the version
```


### Sample report


Command: ```github-digest --cutoff 14 --closed-cutoff 2 coreos/etcd coreos/fleet```.
Private repos are supported, I just picked some active OSS projects.

<table summary="Open pull requests" border="1" width="100%">
    <thead>
    <tr>
        <td>Open PR</td>
        <td>Title</td>
        <td>Size</td>
        <td>Author</td>
        <td>Created</td>
        <td>Updated</td>
    </tr>
    </thead>
    <tbody>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3726">coreos/etcd
                #3726</a>
        </td>
        <td>storage: add store field in watchableStore</td>
        <td>+6/-2 (1)</td>
        <td>yichengq</td>
        <td>2015-10-20 23:13:53 &#43;0000 UTC</td>
        <td>2015-10-20 23:13:55 &#43;0000 UTC</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3725">coreos/etcd
                #3725</a>
        </td>
        <td>Documentation: Fix heading hierarchy.</td>
        <td>+82/-79 (17)</td>
        <td>joshix</td>
        <td>2015-10-20 22:33:08 &#43;0000 UTC</td>
        <td>2015-10-20 22:33:09 &#43;0000 UTC</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3588">coreos/etcd
                #3588</a>
        </td>
        <td>storage/watchable_store.go: use map for unsynced</td>
        <td>+81/-21 (2)</td>
        <td>gyuho</td>
        <td>2015-09-24 10:48:01 &#43;0000 UTC</td>
        <td>2015-10-20 22:09:27 &#43;0000 UTC</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3719">coreos/etcd
                #3719</a>
        </td>
        <td>etcdctl: improve `etcdctl cluster-health`</td>
        <td>+48/-38 (1)</td>
        <td>mqliang</td>
        <td>2015-10-20 03:41:14 &#43;0000 UTC</td>
        <td>2015-10-20 11:10:04 &#43;0000 UTC</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3389">coreos/etcd
                #3389</a>
        </td>
        <td>etcdserver/auth: use fast password gen and check when testing</td>
        <td>+59/-5 (3)</td>
        <td>yichengq</td>
        <td>2015-08-27 22:28:34 &#43;0000 UTC</td>
        <td>2015-10-20 07:16:24 &#43;0000 UTC</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3700">coreos/etcd
                #3700</a>
        </td>
        <td>Replace Summary with Histogram for all metrics</td>
        <td>+46/-37 (8)</td>
        <td>xiang90</td>
        <td>2015-10-17 20:08:09 &#43;0000 UTC</td>
        <td>2015-10-20 06:59:11 &#43;0000 UTC</td>
    </tr>

    </tbody>
</table>


<br/>


<table summary="Closed pull requests" border="1" width="100%">
    <thead>
    <tr>
        <td>Closed PR</td>
        <td>Title</td>
        <td>Author</td>
        <td>Approver</td>
    </tr>
    </thead>
    <tbody>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3727">coreos/etcd
                #3727</a>
        </td>
        <td>raft: fix malformed example name</td>
        <td>yichengq</td>
        <td>yichengq</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3724">coreos/etcd
                #3724</a>
        </td>
        <td>README: fix language for release binaries</td>
        <td>philips</td>
        <td>xiang90</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3720">coreos/etcd
                #3720</a>
        </td>
        <td>rafthttp: deprecate streamTypeMsgApp and remove msgApp stream sent restriction due to streamTypeMsgApp</td>
        <td>yichengq</td>
        <td>yichengq</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3683">coreos/etcd
                #3683</a>
        </td>
        <td>etcdserver: fix raft state machine may block</td>
        <td>yichengq</td>
        <td>yichengq</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3721">coreos/etcd
                #3721</a>
        </td>
        <td>etcdserver: don&#39;t allow methods other than GET in /debug/vars</td>
        <td>mitake</td>
        <td>xiang90</td>
    </tr>

    <tr>
        <td>
            <a href="https://github.com/coreos/etcd/pull/3656">coreos/etcd
                #3656</a>
        </td>
        <td>Added example on how to get node&#39;s value</td>
        <td>kayrus</td>
        <td>yichengq</td>
    </tr>

    </tbody>
</table>


<br/>

<table summary="User stats" border="1" width="50%">
    <thead>
    <tr>
        <td>Member</td>
        <td>Opened</td>
        <td>Closed</td>
        <td>Comments</td>
    </tr>
    </thead>
    <tbody>

    <tr>
        <td>AdoHe</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>MSamman</td>
        <td>1</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>Winslett</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>aheckmann</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>aledbf</td>
        <td>0</td>
        <td>0</td>
        <td>4</td>
    </tr>

    <tr>
        <td>ark76r</td>
        <td>1</td>
        <td>0</td>
        <td>10</td>
    </tr>

    <tr>
        <td>barakmich</td>
        <td>0</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>bcwaldon</td>
        <td>0</td>
        <td>0</td>
        <td>73</td>
    </tr>

    <tr>
        <td>brk0v</td>
        <td>1</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>carmstrong</td>
        <td>0</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>chai2010</td>
        <td>1</td>
        <td>0</td>
        <td>5</td>
    </tr>

    <tr>
        <td>crawford</td>
        <td>1</td>
        <td>0</td>
        <td>11</td>
    </tr>

    <tr>
        <td>divideandconquer</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>ecnahc515</td>
        <td>0</td>
        <td>0</td>
        <td>5</td>
    </tr>

    <tr>
        <td>eparis</td>
        <td>1</td>
        <td>0</td>
        <td>10</td>
    </tr>

    <tr>
        <td>epipho</td>
        <td>1</td>
        <td>0</td>
        <td>21</td>
    </tr>

    <tr>
        <td>gyuho</td>
        <td>3</td>
        <td>0</td>
        <td>39</td>
    </tr>

    <tr>
        <td>hhoover</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>hroyrh</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>iamruinous</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>jonboulle</td>
        <td>4</td>
        <td>3</td>
        <td>49</td>
    </tr>

    <tr>
        <td>joshix</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>judwhite</td>
        <td>0</td>
        <td>0</td>
        <td>8</td>
    </tr>

    <tr>
        <td>junxu</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>kayrus</td>
        <td>1</td>
        <td>0</td>
        <td>4</td>
    </tr>

    <tr>
        <td>kelseyhightower</td>
        <td>0</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>krancour</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>mickep76</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>miekg</td>
        <td>1</td>
        <td>0</td>
        <td>7</td>
    </tr>

    <tr>
        <td>mischief</td>
        <td>0</td>
        <td>0</td>
        <td>6</td>
    </tr>

    <tr>
        <td>mitake</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>mqliang</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>mwitkow-io</td>
        <td>1</td>
        <td>0</td>
        <td>5</td>
    </tr>

    <tr>
        <td>paveltiunov</td>
        <td>1</td>
        <td>0</td>
        <td>0</td>
    </tr>

    <tr>
        <td>philips</td>
        <td>2</td>
        <td>0</td>
        <td>4</td>
    </tr>

    <tr>
        <td>polvi</td>
        <td>1</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>pwhiteside</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>raoofm</td>
        <td>0</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>sbward</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>scole-scea</td>
        <td>0</td>
        <td>0</td>
        <td>5</td>
    </tr>

    <tr>
        <td>smothiki</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>spacejam</td>
        <td>1</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>stuart-warren</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>thepwagner</td>
        <td>1</td>
        <td>0</td>
        <td>2</td>
    </tr>

    <tr>
        <td>umiller</td>
        <td>0</td>
        <td>0</td>
        <td>1</td>
    </tr>

    <tr>
        <td>wuqixuan</td>
        <td>12</td>
        <td>0</td>
        <td>25</td>
    </tr>

    <tr>
        <td>xiang90</td>
        <td>6</td>
        <td>2</td>
        <td>82</td>
    </tr>

    <tr>
        <td>yichengq</td>
        <td>8</td>
        <td>4</td>
        <td>95</td>
    </tr>

    </tbody>
</table>

