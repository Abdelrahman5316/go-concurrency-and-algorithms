

<head>
    <meta charset="UTF-8">
    <title>Go Concurrency and Algorithms</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 40px;
            line-height: 1.6;
        }
        h1, h2 {
            color: #2c3e50;
        }
        nav {
            margin-bottom: 30px;
            padding: 15px;
            background-color: #f4f4f4;
            border-radius: 6px;
        }
        nav a {
            margin-right: 15px;
            text-decoration: none;
            font-weight: bold;
            color: #2980b9;
        }
        nav a:hover {
            text-decoration: underline;
        }
        code {
            background-color: #eeeeee;
            padding: 2px 6px;
            border-radius: 4px;
            font-family: Consolas, monospace;
        }
        .section {
            margin-bottom: 40px;
        }
        footer {
            margin-top: 50px;
            font-size: 14px;
            color: #666;
        }
    </style>
</head>
<body>

<h1>Go Concurrency and Algorithms</h1>

<nav>
    <a href="#overview">Overview</a>
    <a href="#structure">Project Structure</a>
    <a href="#tasks">Implemented Tasks</a>
    <a href="#requirements">Requirements</a>
    <a href="#tests">Running Tests</a>
</nav>

<div class="section" id="overview">
    <h2>Overview</h2>
    <p>
        This repository contains a Go package implementing mathematical computation,
        Unicode-safe string processing, data aggregation, and parallel counting
        using goroutines and channels.
    </p>
</div>

<div class="section" id="structure">
    <h2>Project Structure</h2>
    <ul>
        <li><code>asg1.go</code> – Implementation of tasks</li>
        <li><code>asg1_test.go</code> – Unit tests</li>
        <li><code>list.txt</code> – Input data file for concurrency test</li>
        <li><code>go.mod</code> – Module definition</li>
    </ul>
</div>

<div class="section" id="tasks">
    <h2>Implemented Tasks</h2>

    <h3>1. Sum of Squares</h3>
    <p>Computes <code>0^2 + 1^2 + ... + |n|^2</code></p>
    <p><code>getSumSquares(n int) int</code></p>

    <h3>2. Word Extraction</h3>
    <p>Extracts words ending with a specified rune (Unicode-safe).</p>
    <p><code>getWords(text string, endLetter rune) []string</code></p>

    <h3>3. Course Registration Aggregation</h3>
    <p>Counts unique student registrations per course.</p>
    <p><code>getCourseInfo(records []RegRecord) map[string]int</code></p>

    <h3>4. Parallel Count</h3>
    <p>Counts occurrences of a key using goroutines and channels.</p>
    <p><code>count(list []int, key int, numThreads int) int</code></p>
</div>

<div class="section" id="requirements">
    <h2>Requirements</h2>
    <p>Go 1.18 or newer</p>
    <p>Check installation:</p>
    <p><code>go version</code></p>
</div>

<div class="section" id="tests">
    <h2>Running Tests</h2>
    <p>Execute from project directory:</p>
    <p><code>go test</code></p>
    <p>Verbose mode:</p>
    <p><code>go test -v</code></p>
</div>

<footer>
    <p>Demonstrates Go concurrency, maps, Unicode processing, and unit testing.</p>
</footer>

</body>

