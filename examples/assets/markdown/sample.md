# Example Alert/Callouts

Some text with a GitHub **Important** alert:

This is a standard GitHub-style Alert (also used for non-folding Obsidian Callouts):

```markdown
> [!IMPORTANT]
> This is a GitHub important alert!
```

> [!IMPORTANT]
> This is a GitHub important alert!

---

This is a **Tip** alert using a foldable callout that is closed by default:

```markdown
> [!TIP]-
> This is a GitHub tip in a closed callout.
```

> [!TIP]-
> This is a GitHub tip in a closed callout.
>
> With multiple lines of text.

---

This is an **Info** alert using a foldable callout that is open by default:

```markdown
> [!INFO]+
> This is a an info alert in a foldable callout (open by default).
```

> [!INFO]+
> This is a an info alert in a foldable callout (open by default).

---

This is a **TL;DR** Callout:

```markdown
> [!tldr]
>
> This is a TL;DR (To Long; Didn't Read) callout.
```

> [!tldr]
>
> This is a TL;DR (To Long; Didn't Read) callout.

*...now using the custom title feature...*

```markdown
> [!tldr] TL;DR
>
> This is a TL;DR (To Long; Didn't Read) callout.
```

> [!tldr] TL;DR
>
> This is a TL;DR (To Long; Didn't Read) callout.

---

You can use the custom title to set any title and use an existing recognized marker to pick the icon to use:

```markdown
> [!Important] TL;DR
>
> This is a TL;DR (To Long; Didn't Read) callout, but using the IMPORTANT marker so it uses the color and icon of IMPORTANT.
```

> [!Important] TL;DR
>
> This is a TL;DR (To Long; Didn't Read) callout, but using the IMPORTANT marker so it uses the color and icon of IMPORTANT.

---

```markdown
> [!unknown]
>
> This callout uses an 'unknown' marker name. It defaults to using the "note" icon.
```

> [!unknown]
>
> This callout uses an 'unknown' marker name. It defaults to using the "note" icon.

---

Using `[!noicon]` to create a callout without an icon. By default, this will use the "note", "info" or "default" icon definition (in that order). If none of these three icon kinds have been set, no icon will be used AND

```markdown
> [!noicon]
>
> This callout will not use an icon but will use the marker text 'Noicon' for the header title.
```

> [!noicon]
>
> This callout will not use an icon but will use the marker text 'Noicon' for the header title.

...using a custom title "Wisdom" (still styled as a "note")

```markdown
> [!noicon] Wisdom
>
> This callout will not use an icon but will use the marker text 'Wisdom' for the header title.
```

> [!noicon] Wisdom
>
> This callout will not use an icon but will use the marker text 'Wisdom' for the header title.
