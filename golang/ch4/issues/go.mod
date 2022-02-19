module issues

require (
        github v0.0.0
        issueshtml v0.0.0
)

replace (
        github v0.0.0 => ../github
        issueshtml v0.0.0 => ../issueshtml
)

go 1.17
