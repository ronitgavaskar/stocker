// Package tool is Stocker's anti-corruption layer over external APIs. Each
// tool calls an API and classifies any failure into a fixed ErrorKind,
// returning clean data or a *ToolError. It is a leaf package — depends on
// nothing — so the messy outside world stops here.
package tool
