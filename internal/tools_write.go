package internal

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteTools(s *server.MCPServer, node *Node) {
	// ── Write — Create ───────────────────────────────────────────────────

	s.AddTool(mcp.NewTool("create_frame",
		mcp.WithDescription("Create a new frame on the current page or inside a parent node."),
		mcp.WithNumber("x", mcp.Description("X position (default 0)")),
		mcp.WithNumber("y", mcp.Description("Y position (default 0)")),
		mcp.WithNumber("width", mcp.Description("Width in pixels (default 100)")),
		mcp.WithNumber("height", mcp.Description("Height in pixels (default 100)")),
		mcp.WithString("name", mcp.Description("Frame name")),
		mcp.WithString("fillColor", mcp.Description("Fill color as hex e.g. #FFFFFF")),
		mcp.WithString("layoutMode", mcp.Description("Auto-layout direction: HORIZONTAL, VERTICAL, or NONE")),
		mcp.WithNumber("paddingTop", mcp.Description("Auto-layout top padding")),
		mcp.WithNumber("paddingRight", mcp.Description("Auto-layout right padding")),
		mcp.WithNumber("paddingBottom", mcp.Description("Auto-layout bottom padding")),
		mcp.WithNumber("paddingLeft", mcp.Description("Auto-layout left padding")),
		mcp.WithNumber("itemSpacing", mcp.Description("Auto-layout gap between children")),
		mcp.WithString("primaryAxisAlignItems", mcp.Description("Main-axis alignment: MIN, CENTER, MAX, or SPACE_BETWEEN")),
		mcp.WithString("counterAxisAlignItems", mcp.Description("Cross-axis alignment: MIN, CENTER, MAX, or BASELINE")),
		mcp.WithString("primaryAxisSizingMode", mcp.Description("Main-axis sizing: FIXED or AUTO (hug)")),
		mcp.WithString("counterAxisSizingMode", mcp.Description("Cross-axis sizing: FIXED or AUTO (hug)")),
		mcp.WithString("layoutWrap", mcp.Description("Wrap behaviour: NO_WRAP or WRAP")),
		mcp.WithNumber("counterAxisSpacing", mcp.Description("Gap between wrapped rows/columns (only when layoutWrap is WRAP)")),
		mcp.WithString("parentId", mcp.Description("Parent node ID in colon format. Defaults to current page.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_frame", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_rectangle",
		mcp.WithDescription("Create a new rectangle on the current page or inside a parent node."),
		mcp.WithNumber("x", mcp.Description("X position (default 0)")),
		mcp.WithNumber("y", mcp.Description("Y position (default 0)")),
		mcp.WithNumber("width", mcp.Description("Width in pixels (default 100)")),
		mcp.WithNumber("height", mcp.Description("Height in pixels (default 100)")),
		mcp.WithString("name", mcp.Description("Rectangle name")),
		mcp.WithString("fillColor", mcp.Description("Fill color as hex e.g. #FF5733")),
		mcp.WithNumber("cornerRadius", mcp.Description("Corner radius in pixels")),
		mcp.WithString("parentId", mcp.Description("Parent node ID in colon format. Defaults to current page.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_rectangle", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_ellipse",
		mcp.WithDescription("Create a new ellipse (circle/oval) on the current page or inside a parent node."),
		mcp.WithNumber("x", mcp.Description("X position (default 0)")),
		mcp.WithNumber("y", mcp.Description("Y position (default 0)")),
		mcp.WithNumber("width", mcp.Description("Width in pixels (default 100)")),
		mcp.WithNumber("height", mcp.Description("Height in pixels (default 100)")),
		mcp.WithString("name", mcp.Description("Ellipse name")),
		mcp.WithString("fillColor", mcp.Description("Fill color as hex e.g. #3B82F6")),
		mcp.WithString("parentId", mcp.Description("Parent node ID in colon format. Defaults to current page.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_ellipse", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_text",
		mcp.WithDescription("Create a new text node on the current page or inside a parent node."),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text content"),
		),
		mcp.WithNumber("x", mcp.Description("X position (default 0)")),
		mcp.WithNumber("y", mcp.Description("Y position (default 0)")),
		mcp.WithNumber("fontSize", mcp.Description("Font size in pixels (default 14)")),
		mcp.WithString("fontFamily", mcp.Description("Font family e.g. Inter (default Inter)")),
		mcp.WithString("fontStyle", mcp.Description("Font style e.g. Regular, Bold (default Regular)")),
		mcp.WithString("fillColor", mcp.Description("Text color as hex e.g. #000000")),
		mcp.WithString("name", mcp.Description("Node name")),
		mcp.WithString("parentId", mcp.Description("Parent node ID in colon format. Defaults to current page.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_text", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("import_image",
		mcp.WithDescription("Import a base64-encoded image into Figma as a rectangle with an image fill. Use get_screenshot to capture images or provide your own base64 PNG/JPG."),
		mcp.WithString("imageData",
			mcp.Required(),
			mcp.Description("Base64-encoded image data (PNG or JPG)"),
		),
		mcp.WithNumber("x", mcp.Description("X position (default 0)")),
		mcp.WithNumber("y", mcp.Description("Y position (default 0)")),
		mcp.WithNumber("width", mcp.Description("Width in pixels (default 200)")),
		mcp.WithNumber("height", mcp.Description("Height in pixels (default 200)")),
		mcp.WithString("name", mcp.Description("Node name")),
		mcp.WithString("scaleMode", mcp.Description("Image scale mode: FILL (default), FIT, CROP, or TILE")),
		mcp.WithString("parentId", mcp.Description("Parent node ID in colon format. Defaults to current page.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := map[string]interface{}{
			"imageData": req.GetArguments()["imageData"],
		}
		if x, ok := req.GetArguments()["x"].(float64); ok {
			params["x"] = x
		}
		if y, ok := req.GetArguments()["y"].(float64); ok {
			params["y"] = y
		}
		if w, ok := req.GetArguments()["width"].(float64); ok {
			params["width"] = w
		}
		if h, ok := req.GetArguments()["height"].(float64); ok {
			params["height"] = h
		}
		if n, ok := req.GetArguments()["name"].(string); ok && n != "" {
			params["name"] = n
		}
		if sm, ok := req.GetArguments()["scaleMode"].(string); ok && sm != "" {
			params["scaleMode"] = sm
		}
		if pid, ok := req.GetArguments()["parentId"].(string); ok && pid != "" {
			params["parentId"] = pid
		}
		resp, err := node.Send(ctx, "import_image", nil, params)
		return renderResponse(resp, err)
	})

	// ── Write — Modify ───────────────────────────────────────────────────

	s.AddTool(mcp.NewTool("set_text",
		mcp.WithDescription("Update the text content of an existing TEXT node."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("TEXT node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("New text content"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		text, _ := req.GetArguments()["text"].(string)
		resp, err := node.Send(ctx, "set_text", []string{nodeID}, map[string]interface{}{"text": text})
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("set_fills",
		mcp.WithDescription("Set the fill color of a node."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithString("color",
			mcp.Required(),
			mcp.Description("Fill color as hex e.g. #FF5733"),
		),
		mcp.WithNumber("opacity", mcp.Description("Fill opacity 0–1 (default 1)")),
		mcp.WithString("mode", mcp.Description("'replace' (default) overwrites all fills; 'append' stacks on top of existing fills")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		params := map[string]interface{}{
			"color": req.GetArguments()["color"],
		}
		if op, ok := req.GetArguments()["opacity"].(float64); ok {
			params["opacity"] = op
		}
		if m, ok := req.GetArguments()["mode"].(string); ok {
			params["mode"] = m
		}
		resp, err := node.Send(ctx, "set_fills", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("set_strokes",
		mcp.WithDescription("Set the stroke color and weight of a node."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithString("color",
			mcp.Required(),
			mcp.Description("Stroke color as hex e.g. #000000"),
		),
		mcp.WithNumber("strokeWeight", mcp.Description("Stroke weight in pixels (default 1)")),
		mcp.WithString("mode", mcp.Description("'replace' (default) overwrites all strokes; 'append' stacks on top of existing strokes")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		params := map[string]interface{}{
			"color": req.GetArguments()["color"],
		}
		if sw, ok := req.GetArguments()["strokeWeight"].(float64); ok {
			params["strokeWeight"] = sw
		}
		if m, ok := req.GetArguments()["mode"].(string); ok {
			params["mode"] = m
		}
		resp, err := node.Send(ctx, "set_strokes", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("move_nodes",
		mcp.WithDescription("Move one or more nodes to an absolute position on the canvas."),
		mcp.WithArray("nodeIds",
			mcp.Required(),
			mcp.Description("Node IDs in colon format e.g. ['4029:12345']"),
			mcp.WithStringItems(),
		),
		mcp.WithNumber("x", mcp.Description("Target X position")),
		mcp.WithNumber("y", mcp.Description("Target Y position")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		raw, _ := req.GetArguments()["nodeIds"].([]interface{})
		nodeIDs := toStringSlice(raw)
		params := map[string]interface{}{}
		if x, ok := req.GetArguments()["x"].(float64); ok {
			params["x"] = x
		}
		if y, ok := req.GetArguments()["y"].(float64); ok {
			params["y"] = y
		}
		resp, err := node.Send(ctx, "move_nodes", nodeIDs, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("resize_nodes",
		mcp.WithDescription("Resize one or more nodes. Provide width, height, or both."),
		mcp.WithArray("nodeIds",
			mcp.Required(),
			mcp.Description("Node IDs in colon format e.g. ['4029:12345']"),
			mcp.WithStringItems(),
		),
		mcp.WithNumber("width", mcp.Description("New width in pixels")),
		mcp.WithNumber("height", mcp.Description("New height in pixels")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		raw, _ := req.GetArguments()["nodeIds"].([]interface{})
		nodeIDs := toStringSlice(raw)
		params := map[string]interface{}{}
		if w, ok := req.GetArguments()["width"].(float64); ok {
			params["width"] = w
		}
		if h, ok := req.GetArguments()["height"].(float64); ok {
			params["height"] = h
		}
		resp, err := node.Send(ctx, "resize_nodes", nodeIDs, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("rename_node",
		mcp.WithDescription("Rename a node."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("New name for the node"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		name, _ := req.GetArguments()["name"].(string)
		resp, err := node.Send(ctx, "rename_node", []string{nodeID}, map[string]interface{}{"name": name})
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("clone_node",
		mcp.WithDescription("Clone an existing node, optionally repositioning it or placing it in a new parent."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Source node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithNumber("x", mcp.Description("X position of the clone")),
		mcp.WithNumber("y", mcp.Description("Y position of the clone")),
		mcp.WithString("parentId", mcp.Description("Parent node ID for the clone. Defaults to same parent as source.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		params := map[string]interface{}{}
		if x, ok := req.GetArguments()["x"].(float64); ok {
			params["x"] = x
		}
		if y, ok := req.GetArguments()["y"].(float64); ok {
			params["y"] = y
		}
		if pid, ok := req.GetArguments()["parentId"].(string); ok && pid != "" {
			params["parentId"] = pid
		}
		resp, err := node.Send(ctx, "clone_node", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("set_auto_layout",
		mcp.WithDescription("Set or update auto-layout (flex) properties on an existing frame."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Frame node ID in colon format e.g. '4029:12345'"),
		),
		mcp.WithString("layoutMode", mcp.Description("Auto-layout direction: HORIZONTAL, VERTICAL, or NONE")),
		mcp.WithNumber("paddingTop", mcp.Description("Top padding")),
		mcp.WithNumber("paddingRight", mcp.Description("Right padding")),
		mcp.WithNumber("paddingBottom", mcp.Description("Bottom padding")),
		mcp.WithNumber("paddingLeft", mcp.Description("Left padding")),
		mcp.WithNumber("itemSpacing", mcp.Description("Gap between children")),
		mcp.WithString("primaryAxisAlignItems", mcp.Description("Main-axis alignment: MIN, CENTER, MAX, or SPACE_BETWEEN")),
		mcp.WithString("counterAxisAlignItems", mcp.Description("Cross-axis alignment: MIN, CENTER, MAX, or BASELINE")),
		mcp.WithString("primaryAxisSizingMode", mcp.Description("Main-axis sizing: FIXED or AUTO (hug)")),
		mcp.WithString("counterAxisSizingMode", mcp.Description("Cross-axis sizing: FIXED or AUTO (hug)")),
		mcp.WithString("layoutWrap", mcp.Description("Wrap behaviour: NO_WRAP or WRAP")),
		mcp.WithNumber("counterAxisSpacing", mcp.Description("Gap between wrapped rows/columns (only when layoutWrap is WRAP)")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodeID, _ := req.GetArguments()["nodeId"].(string)
		params := req.GetArguments()
		resp, err := node.Send(ctx, "set_auto_layout", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	// ── Write — Styles ───────────────────────────────────────────────────

	s.AddTool(mcp.NewTool("create_paint_style",
		mcp.WithDescription("Create a new local paint style with a solid fill color."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Style name e.g. 'Brand/Primary'"),
		),
		mcp.WithString("color",
			mcp.Required(),
			mcp.Description("Fill color as hex e.g. #FF5733"),
		),
		mcp.WithString("description", mcp.Description("Optional style description")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_paint_style", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_text_style",
		mcp.WithDescription("Create a new local text style."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Style name e.g. 'Heading/H1'"),
		),
		mcp.WithNumber("fontSize", mcp.Description("Font size in pixels (default 16)")),
		mcp.WithString("fontFamily", mcp.Description("Font family e.g. Inter (default Inter)")),
		mcp.WithString("fontStyle", mcp.Description("Font style e.g. Regular, Bold (default Regular)")),
		mcp.WithString("textDecoration", mcp.Description("NONE, UNDERLINE, or STRIKETHROUGH")),
		mcp.WithNumber("lineHeightValue", mcp.Description("Line height value")),
		mcp.WithString("lineHeightUnit", mcp.Description("Line height unit: PIXELS or PERCENT (default PIXELS)")),
		mcp.WithNumber("letterSpacingValue", mcp.Description("Letter spacing value")),
		mcp.WithString("letterSpacingUnit", mcp.Description("Letter spacing unit: PIXELS or PERCENT (default PIXELS)")),
		mcp.WithString("description", mcp.Description("Optional style description")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_text_style", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_effect_style",
		mcp.WithDescription("Create a new local effect style (drop shadow, inner shadow, or blur)."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Style name e.g. 'Shadow/Card'"),
		),
		mcp.WithString("type", mcp.Description("Effect type: DROP_SHADOW (default), INNER_SHADOW, LAYER_BLUR, or BACKGROUND_BLUR")),
		mcp.WithString("color", mcp.Description("Shadow color as hex e.g. #000000 (default #000000, shadows only)")),
		mcp.WithNumber("opacity", mcp.Description("Shadow color opacity 0–1 (default 0.25, shadows only)")),
		mcp.WithNumber("radius", mcp.Description("Blur radius in pixels (default 8 for shadows, 4 for blurs)")),
		mcp.WithNumber("offsetX", mcp.Description("Shadow X offset in pixels (default 0, shadows only)")),
		mcp.WithNumber("offsetY", mcp.Description("Shadow Y offset in pixels (default 4, shadows only)")),
		mcp.WithNumber("spread", mcp.Description("Shadow spread in pixels (default 0, shadows only)")),
		mcp.WithString("description", mcp.Description("Optional style description")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_effect_style", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_grid_style",
		mcp.WithDescription("Create a new local layout grid style."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Style name e.g. 'Grid/Desktop'"),
		),
		mcp.WithString("pattern", mcp.Description("Grid pattern: GRID (default), COLUMNS, or ROWS")),
		mcp.WithNumber("count", mcp.Description("Number of columns or rows (COLUMNS/ROWS only, default 12)")),
		mcp.WithNumber("gutterSize", mcp.Description("Gutter size in pixels (COLUMNS/ROWS only, default 16)")),
		mcp.WithNumber("offset", mcp.Description("Margin/offset in pixels (COLUMNS/ROWS only, default 0)")),
		mcp.WithString("alignment", mcp.Description("Alignment: STRETCH (default), CENTER, MIN, or MAX (COLUMNS/ROWS only)")),
		mcp.WithNumber("sectionSize", mcp.Description("Grid cell size in pixels (GRID only, default 8)")),
		mcp.WithString("color", mcp.Description("Grid line color as hex e.g. #FF0000 (GRID only, default #FF0000)")),
		mcp.WithNumber("opacity", mcp.Description("Grid line opacity 0–1 (GRID only, default 0.1)")),
		mcp.WithString("description", mcp.Description("Optional style description")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_grid_style", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("update_paint_style",
		mcp.WithDescription("Update the name, color, or description of an existing paint style."),
		mcp.WithString("styleId",
			mcp.Required(),
			mcp.Description("Paint style ID"),
		),
		mcp.WithString("name", mcp.Description("New style name")),
		mcp.WithString("color", mcp.Description("New fill color as hex e.g. #FF5733")),
		mcp.WithString("description", mcp.Description("New style description")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "update_paint_style", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("delete_style",
		mcp.WithDescription("Delete a style (paint, text, effect, or grid) by its ID."),
		mcp.WithString("styleId",
			mcp.Required(),
			mcp.Description("Style ID to delete"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "delete_style", nil, params)
		return renderResponse(resp, err)
	})

	// ── Write — Variables ────────────────────────────────────────────────

	s.AddTool(mcp.NewTool("create_variable_collection",
		mcp.WithDescription("Create a new local variable collection."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Collection name"),
		),
		mcp.WithString("initialModeName", mcp.Description("Name for the initial mode (default 'Mode 1')")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_variable_collection", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("add_variable_mode",
		mcp.WithDescription("Add a new mode to an existing variable collection (e.g. Light/Dark, Desktop/Mobile)."),
		mcp.WithString("collectionId",
			mcp.Required(),
			mcp.Description("Variable collection ID"),
		),
		mcp.WithString("modeName",
			mcp.Required(),
			mcp.Description("Name for the new mode"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "add_variable_mode", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("create_variable",
		mcp.WithDescription("Create a new variable in a collection."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Variable name"),
		),
		mcp.WithString("collectionId",
			mcp.Required(),
			mcp.Description("Variable collection ID"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Variable type: COLOR, FLOAT, STRING, or BOOLEAN"),
		),
		mcp.WithString("value", mcp.Description("Initial value for the first mode. COLOR: hex e.g. #FF5733. FLOAT: number e.g. 16. STRING: text. BOOLEAN: true or false.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "create_variable", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("set_variable_value",
		mcp.WithDescription("Set a variable's value for a specific mode."),
		mcp.WithString("variableId",
			mcp.Required(),
			mcp.Description("Variable ID"),
		),
		mcp.WithString("modeId",
			mcp.Required(),
			mcp.Description("Mode ID within the collection"),
		),
		mcp.WithString("value",
			mcp.Required(),
			mcp.Description("Value to set. COLOR: hex e.g. #FF5733. FLOAT: number e.g. 16. STRING: text. BOOLEAN: true or false."),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "set_variable_value", nil, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("delete_variable",
		mcp.WithDescription("Delete a variable or an entire variable collection. Provide either variableId or collectionId."),
		mcp.WithString("variableId", mcp.Description("Variable ID to delete")),
		mcp.WithString("collectionId", mcp.Description("Collection ID to delete (removes all variables in the collection)")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := req.GetArguments()
		resp, err := node.Send(ctx, "delete_variable", nil, params)
		return renderResponse(resp, err)
	})

	// ── Write — Linked styles & variables ───────────────────────────────

	s.AddTool(mcp.NewTool("apply_style_to_node",
		mcp.WithDescription("Apply an existing local style (paint, text, effect, or grid) to a node, linking the node to that style."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Target node ID in colon format e.g. 4029:12345"),
		),
		mcp.WithString("styleId",
			mcp.Required(),
			mcp.Description("Style ID to apply (from get_styles)"),
		),
		mcp.WithString("target", mcp.Description("For paint styles only — apply to 'fill' (default) or 'stroke'")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := req.GetArguments()
		nodeID, _ := args["nodeId"].(string)
		nodeID = NormalizeNodeID(nodeID)
		params := map[string]interface{}{
			"styleId": args["styleId"],
		}
		if t, ok := args["target"]; ok {
			params["target"] = t
		}
		resp, err := node.Send(ctx, "apply_style_to_node", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("bind_variable_to_node",
		mcp.WithDescription("Bind a local variable to a node property, so the property is driven by the variable's value. Use 'fillColor' to bind a COLOR variable to the node's fill color. Use other fields (opacity, width, height, cornerRadius, itemSpacing, paddingTop, paddingRight, paddingBottom, paddingLeft) for FLOAT variables."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("Target node ID in colon format e.g. 4029:12345"),
		),
		mcp.WithString("variableId",
			mcp.Required(),
			mcp.Description("Variable ID to bind (from get_variable_defs)"),
		),
		mcp.WithString("field",
			mcp.Required(),
			mcp.Description("Property to bind: fillColor | opacity | width | height | cornerRadius | itemSpacing | paddingTop | paddingRight | paddingBottom | paddingLeft"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := req.GetArguments()
		nodeID, _ := args["nodeId"].(string)
		nodeID = NormalizeNodeID(nodeID)
		params := map[string]interface{}{
			"variableId": args["variableId"],
			"field":      args["field"],
		}
		resp, err := node.Send(ctx, "bind_variable_to_node", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	// ── Write — Components ───────────────────────────────────────────────

	s.AddTool(mcp.NewTool("swap_component",
		mcp.WithDescription("Swap the main component of an existing INSTANCE node, replacing it with a different component while keeping position and size."),
		mcp.WithString("nodeId",
			mcp.Required(),
			mcp.Description("INSTANCE node ID in colon format e.g. 4029:12345"),
		),
		mcp.WithString("componentId",
			mcp.Required(),
			mcp.Description("Target COMPONENT node ID in colon format (from get_local_components)"),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := req.GetArguments()
		nodeID, _ := args["nodeId"].(string)
		nodeID = NormalizeNodeID(nodeID)
		componentID, _ := args["componentId"].(string)
		componentID = NormalizeNodeID(componentID)
		params := map[string]interface{}{"componentId": componentID}
		resp, err := node.Send(ctx, "swap_component", []string{nodeID}, params)
		return renderResponse(resp, err)
	})

	s.AddTool(mcp.NewTool("detach_instance",
		mcp.WithDescription("Detach one or more component instances, converting them to plain frames. The link to the main component is broken; all visual properties are preserved."),
		mcp.WithArray("nodeIds",
			mcp.Required(),
			mcp.Description("INSTANCE node IDs in colon format e.g. ['4029:12345']"),
			mcp.WithStringItems(),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		raw, _ := req.GetArguments()["nodeIds"].([]interface{})
		nodeIDs := toStringSlice(raw)
		for i, id := range nodeIDs {
			nodeIDs[i] = NormalizeNodeID(id)
		}
		resp, err := node.Send(ctx, "detach_instance", nodeIDs, nil)
		return renderResponse(resp, err)
	})

	// ── Write — Delete ───────────────────────────────────────────────────

	s.AddTool(mcp.NewTool("delete_nodes",
		mcp.WithDescription("Delete one or more nodes. This cannot be undone via MCP — use with care."),
		mcp.WithArray("nodeIds",
			mcp.Required(),
			mcp.Description("Node IDs to delete in colon format e.g. ['4029:12345']"),
			mcp.WithStringItems(),
		),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		raw, _ := req.GetArguments()["nodeIds"].([]interface{})
		nodeIDs := toStringSlice(raw)
		resp, err := node.Send(ctx, "delete_nodes", nodeIDs, nil)
		return renderResponse(resp, err)
	})
}
