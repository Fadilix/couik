# Couik - Issues & Features Roadmap

A curated list of issues to iterate on the typing CLI app step by step.

---

## 🎨 UI/UX Improvements

### Issue #1: Add colored text feedback for correct/incorrect characters
**Priority:** High  
**Labels:** `enhancement`, `ui`

**Description:**  
Currently, the app only prints characters when correct. Enhance the visual feedback by:
- Showing **green** for correctly typed characters
- Showing **red** for incorrectly typed characters
- Displaying the remaining text in a muted/gray color

**Acceptance Criteria:**
- [ ] Use ANSI escape codes or a library like `fatih/color` for colored output
- [ ] Correct characters appear green
- [ ] Wrong keystrokes trigger a red flash or indicator
- [ ] Untyped text appears in gray/dim

---

### Issue #2: Display real-time WPM while typing
**Priority:** Medium  
**Labels:** `enhancement`, `ui`

**Description:**  
Show the current WPM in real-time as the user types, updating after each character or word.

**Acceptance Criteria:**
- [ ] WPM counter visible during typing session
- [ ] Updates dynamically without disrupting the typing flow
- [ ] Consider placement (e.g., top-right corner or below the text)

---

### Issue #3: Add a progress bar/indicator
**Priority:** Low  
**Labels:** `enhancement`, `ui`

**Description:**  
Show a visual progress bar indicating how much of the quote has been completed.

**Acceptance Criteria:**
- [ ] Progress bar updates in real-time
- [ ] Shows percentage completed
- [ ] Clean visual integration with the rest of the UI

---

## 📊 Statistics & Metrics

### Issue #4: Track and display accuracy percentage
**Priority:** High  
**Labels:** `enhancement`, `feature`

**Description:**  
Calculate and display the user's typing accuracy based on correct vs. total keystrokes.

**Acceptance Criteria:**
- [ ] Track total keystrokes (including mistakes)
- [ ] Calculate accuracy: `(correct_chars / total_keystrokes) * 100`
- [ ] Display accuracy at the end along with WPM
- [ ] Add `CalculateAccuracy()` function to `stats.go`

---

### Issue #5: Add raw WPM vs net WPM calculation
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
- **Raw WPM:** Total characters typed / 5 / minutes
- **Net WPM:** Raw WPM - (errors / minutes)

Both should be displayed at the end of a session.

**Acceptance Criteria:**
- [ ] Implement `CalculateRawWPM()` and `CalculateNetWPM()`
- [ ] Display both metrics in the results screen
- [ ] Document the formulas used

---

### Issue #6: Persist stats history to a local file
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
Save typing session results to a JSON or SQLite file so users can track their progress over time.

**Acceptance Criteria:**
- [ ] Create a `~/.couik/history.json` file (or similar)
- [ ] Store: date, quote, WPM, accuracy, duration
- [ ] Add a `--history` flag to view past results

---

## ⌨️ Game Modes & Features

### Issue #7: Add backspace support for correcting mistakes
**Priority:** High  
**Labels:** `enhancement`, `feature`

**Description:**  
Currently, the game only accepts correct characters. Allow users to press backspace to correct mistakes.

**Acceptance Criteria:**
- [ ] Handle backspace key (`0x7F` or `0x08`)
- [ ] Remove the last character from the visual output
- [ ] Decrement the character index appropriately
- [ ] Track backspace usage for accuracy calculation

---

### Issue #8: Add timed mode (e.g., 15s, 30s, 60s)
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
Add a timed mode where users type as much as they can within a set time limit.

**Acceptance Criteria:**
- [ ] Add CLI flags: `--time 15`, `--time 30`, `--time 60`
- [ ] Display countdown timer during the session
- [ ] End session when time runs out
- [ ] Calculate WPM based on words typed in the time limit

---

### Issue #9: Add word mode (type N words)
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
Add a mode where users type a specific number of words instead of a fixed quote.

**Acceptance Criteria:**
- [ ] Add CLI flag: `--words 25`, `--words 50`, `--words 100`
- [ ] Generate random words from a word list
- [ ] Track progress by word count

---

### Issue #10: Add practice mode for specific characters/keys
**Priority:** Low  
**Labels:** `enhancement`, `feature`

**Description:**  
Allow users to practice specific character sets (e.g., numbers, punctuation, programming symbols).

**Acceptance Criteria:**
- [ ] Add flags like `--practice numbers`, `--practice punctuation`
- [ ] Generate text focusing on the specified characters
- [ ] Provide specialized word lists for each mode

---

## 📚 Content & Quotes

### Issue #11: Fetch quotes from an external API
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
Instead of hardcoded quotes, fetch them from an API like [Quotable](https://github.com/lukePeavey/quotable).

**Acceptance Criteria:**
- [ ] Implement API client for fetching quotes
- [ ] Add fallback to local quotes if network fails
- [ ] Cache quotes locally for offline use
- [ ] Add `--online` flag to force fetching from API

---

### Issue #12: Add programming code snippets mode
**Priority:** Medium  
**Labels:** `enhancement`, `feature`

**Description:**  
Add a mode for practicing typing code snippets in various programming languages.

**Acceptance Criteria:**
- [ ] Add code snippets for popular languages (Go, Python, JavaScript, etc.)
- [ ] Add `--code go`, `--code python` flags
- [ ] Handle special characters common in code

---

### Issue #13: Allow custom text input
**Priority:** Low  
**Labels:** `enhancement`, `feature`

**Description:**  
Allow users to provide their own text to type via file or stdin.

**Acceptance Criteria:**
- [ ] Add `--file /path/to/text.txt` flag
- [ ] Add `--text "custom text here"` flag
- [ ] Validate input length and handle edge cases

---

## 🛠️ CLI & Configuration

### Issue #14: Add CLI flags using Cobra or similar
**Priority:** High  
**Labels:** `enhancement`, `refactor`

**Description:**  
Implement proper CLI argument parsing using a library like `spf13/cobra`.

**Acceptance Criteria:**
- [ ] Add `cmd/` directory structure
- [ ] Implement subcommands: `couik play`, `couik history`, `couik config`
- [ ] Add `--help` documentation for all commands
- [ ] Add version flag `--version`

---

### Issue #15: Add configuration file support
**Priority:** Low  
**Labels:** `enhancement`, `feature`

**Description:**  
Allow users to set default preferences in a config file (`~/.couik/config.yaml`).

**Acceptance Criteria:**
- [ ] Implement config file loading using `viper` or similar
- [ ] Support: default mode, theme, time limit, etc.
- [ ] Add `couik config set <key> <value>` command

---

## 🧪 Testing & Quality

### Issue #16: Add unit tests for stats calculations
**Priority:** High  
**Labels:** `testing`, `quality`

**Description:**  
Add comprehensive unit tests for the `stats` package.

**Acceptance Criteria:**
- [ ] Create `stats_test.go`
- [ ] Test `CalculateTypingSpeed()` with various inputs
- [ ] Test edge cases (zero duration, empty string, etc.)
- [ ] Aim for >80% code coverage

---

### Issue #17: Add integration tests for game loop
**Priority:** Medium  
**Labels:** `testing`, `quality`

**Description:**  
Add tests that simulate a full typing session.

**Acceptance Criteria:**
- [ ] Mock stdin input
- [ ] Verify correct WPM and accuracy calculation
- [ ] Test Ctrl+C interrupt handling

---

## 📖 Documentation

### Issue #18: Write comprehensive README
**Priority:** High  
**Labels:** `documentation`

**Description:**  
Create a detailed README with installation instructions, usage examples, and screenshots/GIFs.

**Acceptance Criteria:**
- [ ] Installation instructions (go install, binary download)
- [ ] Usage examples for all modes
- [ ] Add demo GIF/screenshot
- [ ] Add contributing guidelines

---

### Issue #19: Add CHANGELOG
**Priority:** Low  
**Labels:** `documentation`

**Description:**  
Maintain a changelog following [Keep a Changelog](https://keepachangelog.com/) format.

**Acceptance Criteria:**
- [ ] Create `CHANGELOG.md`
- [ ] Document all changes for each version
- [ ] Follow semantic versioning

---

## 🎯 Suggested Implementation Order

1. **Issue #7** - Backspace support (improves core gameplay)
2. **Issue #1** - Colored text feedback (visual improvement)
3. **Issue #4** - Accuracy tracking (essential metric)
4. **Issue #16** - Unit tests (establish quality baseline)
5. **Issue #14** - CLI flags with Cobra (enables all future features)
6. **Issue #2** - Real-time WPM display
7. **Issue #8** - Timed mode
8. **Issue #18** - Comprehensive README
9. **Issue #6** - Stats history persistence
10. Continue with remaining issues based on preference

---

*Last updated: 2026-01-31*
