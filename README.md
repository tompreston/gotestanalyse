# gotestanalyse
Analyse the JSON output from gotestsum, which is in the [test2json output
format](https://pkg.go.dev/cmd/test2json#hdr-Output_Format). Run it with `--post-run-command`.

Right now it detects flaky tests but it might also report test result (and context) to some backend.

Some examples in the makefile:
```
tompreston github.com/tompreston/gotestanalyse % make
gotestsum \
		--post-run-command ./gotestanalyse \
		--jsonfile test-output.log \
		--rerun-fails \
		--packages . \
		-- \
		-count=1
✖  . (124ms)

DONE 4 tests, 2 skipped, 1 failure in 0.282s

✓  . (117ms)

=== Skipped
=== SKIP: . TestAlwaysSkip (0.00s)
    main_test.go:14:

=== SKIP: . TestAlwaysFail (0.00s)
    main_test.go:19:

=== Failed
=== FAIL: . TestRandomFail (0.00s)
    main_test.go:28: TestRandomFail: This test failed

DONE 2 runs, 5 tests, 2 skipped, 1 failure in 0.553s
4 unique tests, 1 pass, 2 skipped, 1 flaky, 0 failure
```

# 💡 Ideas
Report the actioned TestEvents (pass, skip, fail) and detected flaky tests separately.

For example, it's useful to know both:
* Test1 passed, Test2 failed, Test2 passed (re-run, same context), Test3 skipped
* We know for certain Test2 is flaky
