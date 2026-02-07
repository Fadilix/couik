# Scrolling Text Window Implementation

## Problem

When displaying very long text (like dictionary mode with hundreds of words), rendering the entire text on every keystroke causes performance issues:

1. **String concatenation in loops is slow** - Each `+=` creates a new string allocation
2. **Rendering the entire text** when only a small portion is visible is wasteful
3. **The UI becomes unresponsive** as text length increases

## Solution: Character-Based Sliding Window

Instead of rendering all characters, we only render a "window" of characters around the current cursor position.

### How It Works

```
Full text:  [-----------------------------CURSOR-----------------------------]
                                    ↓
Window:                    [---visible portion---]
                              (only this is rendered)
```

### Implementation Details

#### 1. Calculate Window Size

```go
// Calculate approximately 5 lines worth of characters
textWidth := m.TerminalWidth - 50
if textWidth < 40 {
    textWidth = 40
}
charsPerLine := textWidth
visibleChars := charsPerLine * 5  // ~5 lines of text
```

#### 2. Position the Window Around the Cursor

```go
// Keep the cursor roughly 1/3 from the start of the window
// This gives more "look-ahead" to see upcoming text
windowStart := m.Index - (visibleChars / 3)
if windowStart < 0 {
    windowStart = 0
}

windowEnd := windowStart + visibleChars
if windowEnd > len(m.Target) {
    windowEnd = len(m.Target)
    // Adjust start if we hit the end
    windowStart = windowEnd - visibleChars
    if windowStart < 0 {
        windowStart = 0
    }
}
```

#### 3. Only Render Visible Characters

```go
// Use strings.Builder for efficient string concatenation
var textArea strings.Builder
for i := windowStart; i < windowEnd; i++ {
    s := string(m.Target[i])
    switch {
    case i < m.Index:
        // Already typed - show correct/incorrect styling
        if m.Results[i] {
            textArea.WriteString(correctStyle.Render(s))
        } else {
            textArea.WriteString(wrongStyle.Render(s))
        }
    case i == m.Index:
        // Current cursor position - highlighted
        textArea.WriteString(highlightStyle.Render(s))
    default:
        // Not yet typed - pending style
        textArea.WriteString(pendingStyle.Render(s))
    }
}
```

### Performance Improvements

| Aspect | Before | After |
|--------|--------|-------|
| Characters rendered | All (500+) | ~400 (window) |
| String operations | `+=` (slow) | `strings.Builder` (fast) |
| Complexity per keystroke | O(n) where n = total chars | O(w) where w = window size |

### Visual Effect

As the user types:

```
Start:
[The quick brown fox jumps over ...]
 ^cursor

Mid-way (window scrolls):
[...fox jumps over the lazy dog...]
              ^cursor

Near end (window at end):
[...the lazy dog and went home.]
                           ^cursor
```

## Configuration

The window size is controlled by:

```go
visibleChars := charsPerLine * 5  // Change 5 to adjust visible lines
```

And the cursor position within the window:

```go
windowStart := m.Index - (visibleChars / 3)  // 1/3 from start = 2/3 look-ahead
```

Adjust these values to change:
- **Lines visible**: Multiply `charsPerLine` by desired line count
- **Look-ahead ratio**: Change the divisor (3 = 33% behind, 67% ahead)
