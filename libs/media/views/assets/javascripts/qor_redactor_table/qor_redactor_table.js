$R.add("class", "table.component", {
    mixins: ["dom", "component"],
    init: function (app, el) {
        this.app = app;
        // init
        return el && el.cmnt !== undefined ? el : this._init(el);
    },

    // public
    addHead: function () {
        this.removeHead();

        var $thtds = this.$element.find("tr").first().children("td, th");
        var columns = 0;
        $thtds.each(function (thtd) {
            columns += thtd.colSpan;
        });
        //var columns = $thtds.length;
        var $head = $R.dom("<thead>");
        var $row = this._buildRow(columns, "<th>");

        $head.append($row);
        this.$element.prepend($head);
    },
    addRow: function (columns) {
        var $row = this._buildRow(columns);
        this.$element.append($row);

        return $row;
    },
    addRowTo: function (current, type) {
        return this._addRowTo(current, type);
    },
    addColumnTo: function (current, type) {
        var $current = $R.dom(current);
        var $currentRow = $current.closest("tr");
        var $currentCell = $current.closest("td, th");

        var index = 0;
        $currentRow.find("td, th").each(function (node, i) {
            if (node === $currentCell.get()) index = i;
        });

        this.$element.find("tr").each(function (node) {
            var $node = $R.dom(node);
            var origCell = $node.find("td, th").get(index);
            var $origCell = $R.dom(origCell);

            var $td = $origCell.clone();
            $td.html("");

            if (type === "right") $origCell.after($td);
            else $origCell.before($td);
        });
    },
    removeHead: function () {
        var $head = this.$element.find("thead");
        if ($head.length !== 0) $head.remove();
    },
    removeRow: function (current) {
        var $current = $R.dom(current);
        var $currentRow = $current.closest("tr");

        $currentRow.remove();
    },
    removeColumn: function (current) {
        var $current = $R.dom(current);
        var $currentRow = $current.closest("tr");
        var $currentCell = $current.closest("td, th");

        var index = 0;
        $currentRow.find("td, th").each(function (node, i) {
            if (node === $currentCell.get()) index = i;
        });

        this.$element.find("tr").each(function (node) {
            var $node = $R.dom(node);
            var origCell = $node.find("td, th").get(index);
            var $origCell = $R.dom(origCell);

            $origCell.remove();
        });
    },

    // private
    _init: function (el) {
        var wrapper, element;
        if (typeof el !== "undefined") {
            var $node = $R.dom(el);
            var node = $node.get();
            var $figure = $node.closest("figure");
            if ($figure.length !== 0) {
                wrapper = $figure;
                element = $figure.find("table").get();
            } else if (node.tagName === "TABLE") {
                element = node;
            }
        }

        this._buildWrapper(wrapper);
        this._buildElement(element);
        this._initWrapper();
    },
    _addRowTo: function (current, position) {
        var $current = $R.dom(current);
        var $currentRow = $current.closest("tr");
        if ($currentRow.length !== 0) {
            var columns = $currentRow.children("td, th").length;
            var $newRow = this._buildRow(columns);

            $currentRow[position]($newRow);

            return $newRow;
        }
    },
    _buildRow: function (columns, tag) {
        tag = tag || "<td>";

        var $row = $R.dom("<tr>");
        for (var i = 0; i < columns; i++) {
            var $cell = $R.dom(tag);
            $cell.attr("contenteditable", true);

            $row.append($cell);
        }

        return $row;
    },
    _buildElement: function (node) {
        if (node) {
            this.$element = $R.dom(node);
        } else {
            this.$element = $R.dom("<table>");
            this.append(this.$element);
        }
    },
    _buildWrapper: function (node) {
        node = node || "<figure>";

        this.parse(node);
    },
    _initWrapper: function () {
        this.addClass("redactor-component");
        this.attr({
            "data-redactor-type": "table",
            tabindex: "-1",
            contenteditable: false
        });
    }
});
$R.add("plugin", "table", {
    translations: {
        en: {
            table: "Table",
            "insert-table": "Insert table",
            "insert-row-above": "Insert row above",
            "insert-row-below": "Insert row below",
            "insert-column-left": "Insert column left",
            "insert-column-right": "Insert column right",
            "add-head": "Add head",
            "delete-head": "Delete head",
            "delete-column": "Delete column",
            "delete-row": "Delete row",
            "delete-table": "Delete table",
            "set-table-theme": "Set table theme",
            "set-cell-background": "Set cell background",
            "set-cell-font-size": "Set cell font size",
            "merge-cell": "Merge cell"
        }
    },
    modals: {
        setTableThemeModal: `<form action=""></form>`,
        setCellFontSizeModal: `<form action=""></form>`,
        setCellBackgroundModal: `<form action=""></form>`
    },
    onmodal: {
        setTableThemeModal: {
            open: function ($modal, $form) {
                // add customize className for table
                if (
                    !this.opts.setTableThemeModal &&
                    typeof this.opts.tableClassNames != "undefined"
                ) {
                    var tableClassNames = this.opts.tableClassNames.split(";");
                    var optionsHtml = tableClassNames.reduce(function (memo, data) {
                        const classArr = data.split(",");
                        if (classArr && classArr.length === 2) {
                            const classTitle = classArr[1];
                            const classValue = classArr[0];
                            return (
                                memo + `<option value="${classValue}">${classTitle}</option>`
                            );
                        }
                    }, "");
                    optionsHtml = `<option value="">No Theme</option>` + optionsHtml;

                    this.opts.setTableThemeModal = `<div class="form-item"><label>Select Theme</label><select name="theme">${optionsHtml}
          </select></div>`;
                }
                $form.append(this.opts.setTableThemeModal);
            },
            save: function ($modal, $form) {
                var data = $form.getData();
                if (data && data.theme != undefined) {
                    this.app.api("plugin.table.setTableTheme", {
                        classValue: data.theme
                    });
                }
            }
        },
        setCellFontSizeModal: {
            open: function ($modal, $form) {
                if (
                    !this.opts.setCellFontSizeModal &&
                    typeof this.cellFontSizes != "undefined"
                ) {
                    var optionsHtml = this.cellFontSizes.reduce(function (memo, data) {
                        return memo + `<option value="${data}">${data}px</option>`;
                    }, "");
                    optionsHtml =
                        `<option value="">Remove Font Size</option>` + optionsHtml;

                    this.opts.setCellFontSizeModal = `<div class="form-item"><label>Select Size</label><select name="theme">${optionsHtml}
          </select></div>`;
                }
                $form.append(this.opts.setCellFontSizeModal);
            },
            save: function ($modal, $form) {
                var data = $form.getData();
                if (data && data.theme != undefined) {
                    this.app.api("plugin.table.setCellFontSize", {
                        classValue: data.theme
                    });
                }
            }
        },
        setCellBackgroundModal: {
            open: function ($modal, $form) {
                // add customize className for table
                if (
                    !this.opts.setCellBackgroundModal &&
                    typeof this.opts.cellBackgroundNames != "undefined"
                ) {
                    var cellBackgroundNames = this.opts.cellBackgroundNames.split(";");
                    var optionsHtml = cellBackgroundNames.reduce(function (memo, data) {
                        const classArr = data.split(",");
                        if (classArr && classArr.length === 2) {
                            const classTitle = classArr[1];
                            const classValue = classArr[0];
                            return (
                                memo + `<option value="${classValue}">${classTitle}</option>`
                            );
                        }
                    }, "");
                    optionsHtml = `<option value="">No Background</option>` + optionsHtml;

                    this.opts.setCellBackgroundModal = `<div class="form-item"><label>Select Background</label><select name="theme">${optionsHtml}
          </select></div>`;
                }
                $form.append(this.opts.setCellBackgroundModal);
            },
            save: function ($modal, $form) {
                var data = $form.getData();
                if (data && data.theme != undefined) {
                    this.app.api("plugin.table.setCellBackground", {
                        classValue: data.theme
                    });
                }
            }
        }
    },

    init: function (app) {
        this.app = app;
        this.lang = app.lang;
        this.opts = app.opts;
        this.caret = app.caret;
        this.editor = app.editor;
        this.toolbar = app.toolbar;
        this.component = app.component;
        this.inspector = app.inspector;
        this.insertion = app.insertion;
        this.selection = app.selection;
        this.block = app.block;
    },
    // messages
    oncontextbar: function (e, contextbar) {
        var data = this.inspector.parse(e.target);
        if (data.isComponentType("table")) {
            var node = data.getComponent();
            var buttons = {};

            var selectedCells = $R.dom(node).find("[data-active]");
            if (selectedCells.length > 1) {
                buttons["merge-cell"] = {
                    title: this.lang.get("merge-cell"),
                    api: "plugin.table.mergeCell",
                    args: node
                };
            }

            if (Object.keys(buttons).length) {
                contextbar.set(e, node, buttons);
            } else {
                contextbar.close();
            }
        }
    },
    // messages
    ondropdown: {
        table: {
            observe: function (dropdown) {
                this._observeDropdown(dropdown);
            }
        }
    },
    onbottomclick: function () {
        this.insertion.insertToEnd(this.editor.getLastNode(), "table");
    },

    // public
    start: function () {
        var dropdown = {
            observe: "table",
            "insert-table": {
                title: this.lang.get("insert-table"),
                api: "plugin.table.insert"
            },
            "insert-row-above": {
                title: this.lang.get("insert-row-above"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.addRowAbove"
            },
            "insert-row-below": {
                title: this.lang.get("insert-row-below"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.addRowBelow"
            },
            "insert-column-left": {
                title: this.lang.get("insert-column-left"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.addColumnLeft"
            },
            "insert-column-right": {
                title: this.lang.get("insert-column-right"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.addColumnRight"
            },
            "add-head": {
                title: this.lang.get("add-head"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.addHead"
            },
            "delete-head": {
                title: this.lang.get("delete-head"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.deleteHead"
            },
            "delete-column": {
                title: this.lang.get("delete-column"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.deleteColumn"
            },
            "delete-row": {
                title: this.lang.get("delete-row"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.deleteRow"
            },

            "delete-table": {
                title: this.lang.get("delete-table"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.deleteTable"
            },

            "set-cell-font-size": {
                title: this.lang.get("set-cell-font-size"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.setCellFontSizeModal"
            }
        };

        this.opts.tableClassNames =
            this.opts.tableClassNames ||
            "table-asics-blue,ASICS Blue;table-asics-light-blue,ASICS Light Blue;table-asics-light-green,ASICS Light Green;table-asics-coral,ASICS Carol";
        if (typeof this.opts.tableClassNames != "undefined") {
            dropdown["set-table-theme"] = {
                title: this.lang.get("set-table-theme"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.setTableThemeModal"
            };
        }

        this.opts.cellBackgroundNames =
            "table-cell-asics-blue,ASICS Blue;table-cell-asics-light-blue,ASICS Light Blue;table-cell-asics-light-green,ASICS Light Green;table-cell-asics-coral,ASICS Carol";
        if (typeof this.opts.cellBackgroundNames != "undefined") {
            dropdown["set-cell-background"] = {
                title: this.lang.get("set-cell-background"),
                classname: "redactor-table-item-observable",
                api: "plugin.table.setCellBackgroundModal"
            };
        }

        var obj = {
            title: this.lang.get("table")
        };

        var $button = this.toolbar.addButtonBefore("link", "table", obj);
        $button.setIcon('<i class="re-icon-table"></i>');
        $button.setDropdown(dropdown);

        this.minCellWidth = 40;
        var isMouseUp = false,
            isMouseDown = false,
            isCursorResize = false;
        var startElement, $tdParent;
        var $$editor = $(this.app.editor.$editor.nodes[0]);
        $$editor.on("focus", "td,th", function () {
            $$editor
                .find("td[data-active],th[data-active]")
                .removeAttr("data-active");
            $(this).attr("data-active", "");

            startElement = this;
            $tdParent = $(this).closest("thead,tbody");
        });

        $("body").on("click", function (e) {
            var isInTable = $(e.target).closest("table").length > 0;
            var isInToolbar = $(e.target).closest(".redactor-toolbar").length > 0;
            var isInDropDown = $(e.target).closest(".redactor-dropdown").length > 0;
            var isModalShow = $(".redactor-modal.open").length > 0;
            if (!isInTable && !isInToolbar && !isInDropDown && !isModalShow) {
                $$editor
                    .find("td[data-active],th[data-active]")
                    .removeAttr("data-active");
            }
        });

        $$editor.on("mousedown", "td, th", function (e) {
            isCursorResize = $(this).css("cursor") == "col-resize";
            var $currentTable = $(this).closest("table");
            var $tableComponent = $(this).closest(
                ".redactor-component[data-redactor-type='table']"
            );
            _this.calcTableElementPosition($currentTable);

            if (isCursorResize) {
                // show resize
                _this.isColResize = true;
                _this.colResizeStartX = e.clientX;
                _this.colResizeStartElement = e.target;
                _this.colResizeStartPositionX = e.target.offsetLeft + e.offsetX;
                // get init left point position
                e.preventDefault();
                var $tableLine = $('<div class="redactor-component-table-line"></div>');
                $tableLine.css("left", _this.colResizeStartPositionX);
                $tableComponent.append($tableLine);
            } else {
                isMouseUp = false;
                isMouseDown = true;
            }
        });
        $("body").on("mouseup", function () {
            // clear helper element
            $(".redactor-component-table-line").remove();
        });

        $$editor.on(
            "mouseup",
            '.redactor-component[data-redactor-type="table"]',
            function (e) {
                isMouseUp = true;
                isMouseDown = false;
                //console.log("show context bar");
                var $tableComponent = $(this);

                var selectedCells = $tableComponent.find("[data-active]");
                _this.selectedCells = selectedCells;

                if (_this.isColResize) {
                    // get end point and calc with

                    _this.isColResize = false;
                    var changeColIndex = _this.colResizeStartElement.lastPoint.col;
                    _this.renderColGroup(
                        $(this).find("table"),
                        changeColIndex,
                        _this.changeColWidth
                    );
                    $tableComponent.find(".redactor-component-table-line").remove();
                }
            }
        );

        $$editor.on("mousemove", "td, th", function (e) {
            if (e.target.offsetWidth - e.offsetX < 10 && !isMouseDown) {
                $(this).css("cursor", "col-resize");
            } else {
                $(this).css("cursor", "");
            }

            if (_this.isColResize) {
                var colWidth;
                var travelWidth = _this.colResizeStartX - e.clientX;
                var $tableComponent = $(this).closest(
                    ".redactor-component[data-redactor-type='table']"
                );

                var $tableBody = $tableComponent.find("thead,tbody");
                var tableWidth = $tableBody.outerWidth();
                var changeColIndex = _this.colResizeStartElement.lastPoint.col;

                // 记录 当前索引 cell 的 colspan 并计算最小值，最小值不是1 则需要重新计算travel宽度
                var minColSpansLeft = null;
                var minColSpansRight = null;

                for (var i = 0; i < _this.tableHeadElements.length; i++) {
                    var cell = _this.tableHeadElements[i][changeColIndex];
                    minColSpansLeft = minColSpansLeft || cell;
                    if (minColSpansLeft.colSpan < cell.colSpan) {
                        minColSpansLeft = cell;
                    }

                    var nextCell = _this.tableHeadElements[i][changeColIndex + 1];
                    minColSpansRight = minColSpansRight || nextCell;
                    if (minColSpansRight.colSpan < nextCell.colSpan) {
                        minColSpansRight = nextCell;
                    }
                }
                for (var i = 0; i < _this.tableBodyElements.length; i++) {
                    var cell = _this.tableBodyElements[i][changeColIndex];
                    minColSpansLeft = minColSpansLeft || cell;
                    if (minColSpansLeft.colSpan > cell.colSpan) {
                        minColSpansLeft = cell;
                    }

                    var nextCell = _this.tableBodyElements[i][changeColIndex + 1];
                    minColSpansRight = minColSpansRight || nextCell;
                    if (minColSpansRight.colSpan > nextCell.colSpan) {
                        minColSpansRight = nextCell;
                    }
                }

                var $colGroupCols = $tableComponent.find("colgroup").find("col");
                if (travelWidth > 0) {
                    // left current cell
                    var $col = $colGroupCols.eq(changeColIndex);
                    if ($col.length) {
                        colWidth = $col.attr("width").split("%")[0] / 100 * tableWidth;
                        if (minColSpansLeft.colSpan > 1) {
                            colWidth = minColSpansLeft.offsetWidth;
                        }
                    } else {
                        colWidth = e.target.offsetWidth / e.target.colSpan;
                        if (minColSpansLeft.colSpan > 1) {
                            // if this col is take full rows and cols, colWidth should update
                            colWidth = minColSpansLeft.offsetWidth;
                        }
                    }
                } else {
                    // right cell
                    var $col = $colGroupCols.eq(changeColIndex + 1);

                    if ($col.length) {
                        colWidth = $col.attr("width").split("%")[0] / 100 * tableWidth;

                        if (minColSpansRight.colSpan > 1) {
                            colWidth = minColSpansRight.offsetWidth;
                        }
                    } else {
                        colWidth = e.target.offsetWidth / e.target.colSpan;
                        if (minColSpansRight.colSpan > 1) {
                            colWidth = minColSpansRight.offsetWidth;
                        }
                    }
                }

                if (colWidth - _this.minCellWidth >= Math.abs(travelWidth) > 0) {
                    var changeColWidth = e.clientX - _this.colResizeStartX;

                    $tableComponent
                        .find(".redactor-component-table-line")
                        .css("left", _this.colResizeStartPositionX + changeColWidth);

                    _this.changeColWidth = changeColWidth;
                }
            }
        });

        $$editor.on("mouseout", "td, th", function () {
        });

        var _this = this;

        $$editor.on("mouseenter", "td, th", function () {
            var endElement = this;
            var $currentTdParent = $(this).closest("thead,tbody");
            if (isMouseDown && $currentTdParent[0] === $tdParent[0]) {
                $currentTdParent
                    .find("td[data-active],th[data-active]")
                    .removeAttr("data-active");
                _this.calcCellMergeRange(startElement, endElement);
                _this.renderCellMergeRange(this.tagName, _this.finalRange);
            }
        });
    },

    renderColGroup: function ($table, changeColIndex, changeColWidth) {
        var _this = this;
        var $colGroup = $table.find("colgroup");
        var $cols = $colGroup.find("col");
        var tableWidth = $table.find("thead,tbody").outerWidth();
        var minColWidthPercent = (_this.minCellWidth / tableWidth * 100).toFixed(2);
        var minColWidth = tableWidth * minColWidthPercent / 100;

        if (changeColIndex != undefined) {
            // 记录 当前索引 cell 的 colspan 并计算最小值，最小值不是1 则需要重新计算travel宽度
            var minColSpansLeft = null;
            var minColSpansRight = null;

            for (var i = 0; i < _this.tableHeadElements.length; i++) {
                var cell = _this.tableHeadElements[i][changeColIndex];
                minColSpansLeft = minColSpansLeft || cell;
                if (minColSpansLeft.colSpan < cell.colSpan) {
                    minColSpansLeft = cell;
                }

                var nextCell = _this.tableHeadElements[i][changeColIndex + 1];
                minColSpansRight = minColSpansRight || nextCell;
                if (minColSpansRight.colSpan < nextCell.colSpan) {
                    minColSpansRight = nextCell;
                }
            }
            for (var i = 0; i < _this.tableBodyElements.length; i++) {
                var cell = _this.tableBodyElements[i][changeColIndex];
                minColSpansLeft = minColSpansLeft || cell;
                if (minColSpansLeft.colSpan > cell.colSpan) {
                    minColSpansLeft = cell;
                }

                var nextCell = _this.tableBodyElements[i][changeColIndex + 1];
                minColSpansRight = minColSpansRight || nextCell;
                if (minColSpansRight.colSpan > nextCell.colSpan) {
                    minColSpansRight = nextCell;
                }
            }
        }

        if (
            $colGroup.length &&
            $cols.length == _this.tableBodyElements[0].length &&
            changeColIndex != undefined &&
            changeColWidth != undefined
        ) {
            var currentCol = $colGroup.find("col").eq(changeColIndex);
            var nextChangeColIndex = changeColIndex + 1;
            var nextCol = $colGroup.find("col").eq(nextChangeColIndex);

            var currentWidth =
                currentCol.attr("width").split("%")[0] / 100 * tableWidth;
            var currentChangeWidth =
                currentWidth * minColSpansLeft.colSpan + changeColWidth;
            var nextWidth = nextCol.attr("width").split("%")[0] / 100 * tableWidth;
            var nextChangeWidth =
                nextWidth * minColSpansRight.colSpan - changeColWidth;

            // range limit detect

            if (currentChangeWidth < minColWidth) {
                currentChangeWidth = minColWidth;

                changeColWidth = minColWidth - currentWidth * minColSpansLeft.colSpan;
            }

            if (nextChangeWidth < minColWidth) {
                nextChangeWidth = minColWidth;

                changeColWidth = nextWidth * minColSpansRight.colSpan - minColWidth;
            }

            // re calc
            var currentWidthPercent =
                (currentWidth * minColSpansLeft.colSpan + changeColWidth) /
                tableWidth *
                100;
            var nextWidthPercent =
                (nextWidth * minColSpansRight.colSpan - changeColWidth) /
                tableWidth *
                100;

            // support currentCol or nextCol is big cell (more then 2colspan)
            //
            var minColSpanLeftNumber = minColSpansLeft.colSpan;
            var minColSpanRightNumber = minColSpansRight.colSpan;

            while (minColSpanLeftNumber--) {
                var index = changeColIndex - minColSpanLeftNumber;
                $cols
                    .eq(index)
                    .attr(
                        "width",
                        (currentWidthPercent / minColSpansLeft.colSpan).toFixed(2) + "%"
                    );
            }

            while (minColSpanRightNumber--) {
                var index = changeColIndex + minColSpanRightNumber + 1;
                $cols
                    .eq(index)
                    .attr(
                        "width",
                        (nextWidthPercent / minColSpansRight.colSpan).toFixed(2) + "%"
                    );
            }
        } else {
            $colGroup.remove();
            // calc all cell
            var colgroupCol = _this.tableBodyElements[0].reduce(function (
                memo,
                tableElement,
                tableElementIndex
            ) {
                var colspan = $(tableElement).attr("colspan") || 1;
                var width = $(tableElement).outerWidth();
                width = width / colspan;

                //  if bigcell  update changedWidth
                if (changeColIndex != undefined) {
                    if (
                        changeColIndex - minColSpansLeft.colSpan + 1 <= tableElementIndex &&
                        tableElementIndex <= changeColIndex
                    ) {
                        width =
                            (Number(width * minColSpansLeft.colSpan) +
                                Number(changeColWidth)) /
                            minColSpansLeft.colSpan;
                    }
                    if (
                        changeColIndex + 1 <= tableElementIndex &&
                        tableElementIndex <=
                        changeColIndex + 1 + -1 + minColSpansRight.colSpan
                    ) {
                        width =
                            (Number(width * minColSpansRight.colSpan) -
                                Number(changeColWidth)) /
                            minColSpansRight.colSpan;
                    }
                }

                var widthPercent = width / tableWidth * 100;

                return memo + `<col width="${widthPercent.toFixed(2)}%">`;
            }, "");

            var colGroupHtml = `<colgroup>${colgroupCol}</colgroup>`;
            $table.prepend(colGroupHtml);
        }
    },

    calcCellMergeRange: function (startElement, endElement) {
        var _this = this;

        var _tableElements = _this.tableBodyElements;
        if (startElement.tagName === "TH") {
            _tableElements = _this.tableHeadElements;
        }

        var startPoint = {
            firstPoint: startElement.firstPoint,
            lastPoint: startElement.lastPoint
        };
        var endPoint = {
            firstPoint: endElement.firstPoint,
            lastPoint: endElement.lastPoint
        };

        var minRow = Math.min(
            startPoint.firstPoint.row,
            startPoint.lastPoint.row,
            endPoint.firstPoint.row,
            endPoint.lastPoint.row
            ),
            minCol = Math.min(
                startPoint.firstPoint.col,
                startPoint.lastPoint.col,
                endPoint.firstPoint.col,
                endPoint.lastPoint.col
            ),
            maxRow = Math.max(
                startPoint.firstPoint.row,
                startPoint.lastPoint.row,
                endPoint.firstPoint.row,
                endPoint.lastPoint.row
            ),
            maxCol = Math.max(
                startPoint.firstPoint.col,
                startPoint.lastPoint.col,
                endPoint.firstPoint.col,
                endPoint.lastPoint.col
            );

        function loop(minRow, minCol, maxRow, maxCol) {
            var newMinRow = minRow;
            var newMinCol = minCol;
            var newMaxRow = maxRow;
            var newMaxCol = maxCol;
            for (var x = minRow; x <= maxRow; x++) {
                for (var y = minCol; y <= maxCol; y++) {
                    var currentElement = _tableElements[x][y];
                    var currentStartPoint = currentElement.firstPoint;
                    var currentEndPoint = currentElement.lastPoint;

                    newMinRow = Math.min(
                        parseInt(currentStartPoint.row),
                        parseInt(currentEndPoint.row),
                        newMinRow
                    );
                    newMinCol = Math.min(
                        currentStartPoint.col,
                        currentEndPoint.col,
                        newMinCol
                    );
                    newMaxRow = Math.max(
                        currentStartPoint.row,
                        currentEndPoint.row,
                        newMaxRow
                    );
                    newMaxCol = Math.max(
                        currentStartPoint.col,
                        currentEndPoint.col,
                        newMaxCol
                    );
                }
            }

            if (
                newMinRow == minRow &&
                newMinCol == minCol &&
                newMaxRow == maxRow &&
                newMaxCol == maxCol
            ) {
                return {
                    minRow,
                    minCol,
                    maxRow,
                    maxCol
                };
            } else {
                return loop(newMinRow, newMinCol, newMaxRow, newMaxCol);
            }
        }

        _this.finalRange = loop(minRow, minCol, maxRow, maxCol);
        _this.selectedRowRange = maxRow - minRow + 1;
        _this.selectedColRange = maxCol - minCol + 1;
    },

    renderCellMergeRange: function (tagName, {minRow, minCol, maxRow, maxCol}) {
        var _this = this;
        var _tableElements = _this.tableBodyElements;
        if (tagName === "TH") {
            _tableElements = _this.tableHeadElements;
        }

        for (var x = minRow; x <= maxRow; x++) {
            for (var y = minCol; y <= maxCol; y++) {
                var currentElement = _tableElements[x][y];
                $(currentElement).attr("data-active", "");
            }
        }
    },

    savePosition: function (element, point) {
        if (!element.lastPoint) {
            element.firstPoint = point;
        }
        element.lastPoint = point;
    },

    calcTableElementPosition: function ($table) {
        // set style  cursor: text !important; for last td
        var _this = this;
        this.tableBodyElements = [];
        this.tableHeadElements = [];
        var $thead = $table.find("thead"),
            $tbody = $table.find("tbody");

        function fill(element, indexTr, tdIndex, rowSpanLength, colSpanLength) {
            var tagName = element.tagName;
            var tableElement = _this.tableBodyElements;
            if (tagName === "TH") {
                tableElement = _this.tableHeadElements;
            }
            var maxIndexTr = parseInt(indexTr) + parseInt(rowSpanLength);
            var maxTdIndex = parseInt(colSpanLength) + parseInt(tdIndex);

            for (var i = indexTr; i < maxIndexTr; i++) {
                for (var j = tdIndex; j < maxTdIndex; j++) {
                    tableElement[i] = tableElement[i] || [];
                    tableElement[i][j] = element;
                }
            }
        }

        [$thead, $tbody].map(function ($tableTHTB) {
            var _tableElement = _this.tableBodyElements;
            if ($tableTHTB === $thead) {
                _tableElement = _this.tableHeadElements;
            }
            $tableTHTB.find("tr").each(function (indexTr) {
                var $currentRow = $(this);
                _tableElement[indexTr] = _tableElement[indexTr] || [];

                var tdIndex = 0;
                var $findTdThs = $currentRow.find("td,th");
                $findTdThs.each(function () {
                    var rowSpanLength = $(this).attr("rowspan") || 1;
                    var colSpanLength = $(this).attr("colspan") || 1;

                    function getInserCell(indexTr, tdIndex) {
                        var currentElement = _tableElement[indexTr][tdIndex];
                        var currentRowSpanLength = $(currentElement).attr("rowspan") || 1;
                        var currentColSpanLength = $(currentElement).attr("colspan") || 1;
                        if (currentRowSpanLength > 1 || currentColSpanLength > 1) {
                            tdIndex += Number(currentColSpanLength);
                            return getInserCell(indexTr, tdIndex);
                        } else {
                            return {indexTr, tdIndex};
                        }
                    }

                    if (!_tableElement[indexTr][tdIndex]) {
                        if (rowSpanLength > 1 || colSpanLength > 1) {
                            fill(this, indexTr, tdIndex, rowSpanLength, colSpanLength);
                            tdIndex += Number(colSpanLength);
                        } else {
                            _tableElement[indexTr][tdIndex] = this;
                            tdIndex++;
                        }
                    } else {
                        var currentElement = _tableElement[indexTr][tdIndex];
                        var currentRowSpanLength = $(currentElement).attr("rowspan") || 1;
                        var currentColSpanLength = $(currentElement).attr("colspan") || 1;
                        if (currentRowSpanLength > 1 || currentColSpanLength > 1) {
                            var getInsertPoint = getInserCell(indexTr, tdIndex);
                            tdIndex = Number(getInsertPoint.tdIndex);

                            if (rowSpanLength > 1 || colSpanLength > 1) {
                                // 若待插入的元素也是big cell 执行 fill
                                fill(this, indexTr, tdIndex, rowSpanLength, colSpanLength);
                                tdIndex += Number(colSpanLength);
                            } else {
                                _tableElement[indexTr][tdIndex] = this;
                                tdIndex++;
                            }
                        } else {
                            tdIndex++;
                        }
                    }
                });
            });
        });

        [_this.tableBodyElements, _this.tableHeadElements].map(function (
            tableElements
        ) {
            tableElements.forEach(function (tableTr, row) {
                tableTr.forEach(function (tableTd, col) {
                    tableTd.lastPoint = undefined;
                });
            });

            tableElements.forEach(function (tableTr, row) {
                tableTr.forEach(function (tableTd, col) {
                    _this.savePosition(tableTd, {row, col});
                    // debug code
                    // tableTd.innerHTML = `${JSON.stringify(
                    //   tableTd.firstPoint
                    // )},${JSON.stringify(tableTd.lastPoint)}`;
                });
            });
        });
    },
    insert: function () {
        var rows = 2;
        var columns = 3;
        var $component = this.component.create("table");
        $component.$element.addClass("table-asics-richeditor");

        for (var i = 0; i < rows; i++) {
            $component.addRow(columns);
        }

        $component = this.insertion.insertHtml($component);
        this.caret.setStart($component);
    },
    addRowAbove: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");
            var currentTd = $currentTd[0];
            var appendTableTag = `td`;
            var _tableElements = _this.tableBodyElements;
            if (currentTd.tagName === "TH") {
                _tableElements = _this.tableHeadElements;
                appendTableTag = `th`;
            }

            var currentRow = currentTd.firstPoint.row;
            var allColsLength = _tableElements[0].length;

            var $addRow = $("<tr></tr>");

            for (var i = 0; i < allColsLength; i++) {
                var cell = _tableElements[currentRow][i];
                var colspan = parseInt($(cell).attr("colspan") || 1);
                var rowspan = parseInt($(cell).attr("rowspan") || 1);
                if (cell.firstPoint.row == currentRow) {
                    $addRow.append(
                        `<${appendTableTag} contenteditable="true"></${appendTableTag}>`
                    );
                } else {
                    // bigcell  is not on this row
                    if (colspan > 1 || rowspan > 1) {
                        //  big cell
                        if (!$(cell).attr("data-spanupdated")) {
                            if (rowspan > 1) {
                                $(cell).attr("rowspan", rowspan + 1);
                                $(cell).attr("data-spanupdated", "true");
                            }
                        }
                    }
                }
            }
            $addRow.insertBefore($(current).closest("tr"));
            $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
            _this.calcTableElementPosition($(table));
            _this.renderColGroup($(table));
        }
    },
    addRowBelow: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");
            var currentTd = $currentTd[0];
            var _tableElements = _this.tableBodyElements;
            var appendTableTag = `td`;
            if (currentTd.tagName === "TH") {
                _tableElements = _this.tableHeadElements;
                appendTableTag = `th`;
            }
            var currentRow = currentTd.lastPoint.row;
            var allColsLength = _tableElements[0].length;

            var $addRow = $("<tr></tr>");

            for (var i = 0; i < allColsLength; i++) {
                var cell = _tableElements[currentRow][i];
                var colspan = parseInt($(cell).attr("colspan") || 1);
                var rowspan = parseInt($(cell).attr("rowspan") || 1);
                if (cell.lastPoint.row == currentRow) {
                    $addRow.append(
                        `<${appendTableTag} contenteditable="true"></${appendTableTag}>`
                    );
                } else {
                    // bigcell  is not on this row
                    if (colspan > 1 || rowspan > 1) {
                        //  big cell
                        if (!$(cell).attr("data-spanupdated")) {
                            if (rowspan > 1) {
                                $(cell).attr("rowspan", rowspan + 1);
                                $(cell).attr("data-spanupdated", "true");
                            }
                        }
                    }
                }
            }

            // get big cell right insert position
            var targetTdTh = _tableElements[currentTd.lastPoint.row].filter(function (
                element
            ) {
                return element.colSpan == 1 && element.rowSpan == 1;
            });
            if (targetTdTh.length > 0) {
                $addRow.insertAfter($(targetTdTh[0]).closest("tr"));
                $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
                _this.calcTableElementPosition($(table));
                _this.renderColGroup($(table));
            }
        }
    },
    addColumnLeft: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");

            var currentTd = $currentTd[0];
            var currentCol = currentTd.firstPoint.col;

            [_this.tableBodyElements, _this.tableHeadElements].map(function (
                tableElements
            ) {
                var allRowsLength = tableElements.length;
                var insertTag = `td`;
                if (tableElements === _this.tableHeadElements) {
                    insertTag = `th`;
                }

                for (var i = 0; i < allRowsLength; i++) {
                    var cell = tableElements[i][currentCol];
                    var colspan = parseInt($(cell).attr("colspan") || 1);
                    var rowspan = parseInt($(cell).attr("rowspan") || 1);

                    if (currentCol == 0) {
                        $(cell)
                            .closest("thead,tbody")
                            .find("tr")
                            .eq(i)
                            .prepend(`<${insertTag} contenteditable="true"></${insertTag}>`);
                    } else {
                        if (cell.firstPoint.row == i) {
                            if (colspan > 1 || rowspan > 1) {
                                //  big cell
                                if (!$(cell).attr("data-spanupdated")) {
                                    $(cell).attr("colspan", colspan + 1);
                                    $(cell).attr("data-spanupdated", "true");
                                }
                            } else {
                                $(
                                    `<${insertTag} contenteditable="true"></${insertTag}>`
                                ).insertBefore(cell);
                            }
                        }
                    }
                }
            });

            $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
            _this.calcTableElementPosition($(table));
            _this.renderColGroup($(table));
        }
    },
    addColumnRight: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");
            var currentTd = $currentTd[0];
            var currentCol = currentTd.lastPoint.col;

            [_this.tableBodyElements, _this.tableHeadElements].map(function (
                tableElements
            ) {
                var allRowsLength = tableElements.length;
                var insertTag = `td`;
                if (tableElements === _this.tableHeadElements) {
                    insertTag = `th`;
                }

                var allColsLength = tableElements[0] instanceof Array ? tableElements[0].length : 0;

                for (var i = 0; i < allRowsLength; i++) {
                    var cell = tableElements[i][currentCol];
                    var colspan = parseInt($(cell).attr("colspan") || 1);
                    var rowspan = parseInt($(cell).attr("rowspan") || 1);

                    if (currentCol == allColsLength - 1) {
                        $(cell)
                            .closest("thead,tbody")
                            .find("tr")
                            .eq(i)
                            .append(`<${insertTag} contenteditable="true"></${insertTag}>`);
                    } else {
                        if (cell.lastPoint.row == i) {
                            if (colspan > 1 || rowspan > 1) {
                                //  big cell
                                if (!$(cell).attr("data-spanupdated")) {
                                    $(cell).attr("colspan", colspan + 1);
                                    $(cell).attr("data-spanupdated", "true");
                                }
                            } else {
                                $(
                                    `<${insertTag} contenteditable="true"></${insertTag}>`
                                ).insertAfter(cell);
                            }
                        }
                    }
                }
            });

            $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
            _this.calcTableElementPosition($(table));
            _this.renderColGroup($(table));
        }
    },
    addHead: function () {
        var $component = this._getComponent();
        if ($component) {
            this.selection.save();
            $component.addHead();
            this.selection.restore();
        }
    },
    deleteHead: function () {
        var $component = this._getComponent();
        if ($component) {
            var current = this.selection.getCurrent();
            var $head = $R.dom(current).closest("thead");
            if ($head.length !== 0) {
                $component.removeHead();
                this.caret.setStart($component);
            } else {
                this.selection.save();
                $component.removeHead();
                this.selection.restore();
            }
        }
    },
    deleteColumn: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");
            var currentTd = $currentTd[0];
            var currentCol = currentTd.firstPoint.col;

            [_this.tableBodyElements, _this.tableHeadElements].map(function (
                tableElements
            ) {
                var allRowsLength = tableElements.length;

                for (var i = 0; i < allRowsLength; i++) {
                    var cell = tableElements[i][currentCol];
                    var colspan = parseInt($(cell).attr("colspan") || 1);
                    var rowspan = parseInt($(cell).attr("rowspan") || 1);
                    if (cell.firstPoint.row == i) {
                        if (colspan > 1 || rowspan > 1) {
                            //  big cell
                            if (!$(cell).attr("data-spanupdated")) {
                                if (colspan > 1) {
                                    $(cell).attr("colspan", colspan - 1);
                                    $(cell).attr("data-spanupdated", "true");
                                } else {
                                    $(cell).remove();
                                }
                            }
                        } else {
                            $(cell).remove();
                        }
                    }
                }
            });

            $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
            _this.calcTableElementPosition($(table));
            _this.renderColGroup($(table));
        }
    },
    deleteRow: function () {
        // todo big cell
        var _this = this;
        var table = this._getTable();
        if (table) {
            var current = this.selection.getCurrent();
            var $currentTd = $(current).closest("td,th");
            var currentTd = $currentTd[0];
            var _tableElements = _this.tableBodyElements;
            if (currentTd.tagName === "TH") {
                _tableElements = _this.tableHeadElements;
            }
            var $currentTr = $currentTd.closest("tr");
            var currentRow = currentTd.firstPoint.row;
            var allColsLength = _tableElements[0].length;
            for (var i = 0; i < allColsLength; i++) {
                var cell = _tableElements[currentRow][i];
                var colspan = parseInt($(cell).attr("colspan") || 1);
                var rowspan = parseInt($(cell).attr("rowspan") || 1);
                if (cell.firstPoint.col == i) {
                    if (colspan > 1 || rowspan > 1) {
                        //  big cell
                        if (!$(cell).attr("data-spanupdated")) {
                            if (rowspan > 1) {
                                $(cell).attr("rowspan", rowspan - 1);
                                $(cell).attr("data-spanupdated", "true");

                                if (cell.firstPoint.row == currentRow) {
                                    //  move to next line
                                    if (i - 1 < 0) {
                                        $currentTr.next().prepend(cell);
                                    } else {
                                        $(cell).insertAfter(_tableElements[currentRow + 1][i - 1]);
                                    }
                                }
                            }
                        } else {
                            $(cell).remove();
                        }
                    } else {
                        // may be is th
                        $(cell).remove();
                    }
                }
            }
            $currentTr.remove();

            $(table).find("[data-spanupdated]").removeAttr("data-spanupdated");
            _this.calcTableElementPosition($(table));
            _this.renderColGroup($(table));
        }
    },
    mergeCell: function () {
        var _this = this;
        var table = this._getTable();
        if (table) {
            var $component = this._getComponent();
            if ($component) {
                var current = this.selection.getCurrent();

                var $tdParent = $(current).closest("tbody,thead");

                var _tableElements = _this.tableBodyElements;
                if ($tdParent[0].tagName === "THEAD") {
                    _tableElements = _this.tableHeadElements;
                }

                var {minRow, minCol, maxRow, maxCol} = _this.finalRange;
                var $firstCell = $(_tableElements[minRow][minCol]);
                var rowSpanLength = $firstCell.attr("rowspan") || 1;
                var colSpanLength = $firstCell.attr("colspan") || 1;

                rowSpanLength = Math.max(rowSpanLength - 1, _this.selectedRowRange);
                $firstCell.attr("rowspan", rowSpanLength);
                colSpanLength = Math.max(colSpanLength - 1, _this.selectedColRange);
                $firstCell.attr("colspan", colSpanLength);

                for (var x = minRow; x <= maxRow; x++) {
                    for (var y = minCol; y <= maxCol; y++) {
                        var currentElement = _tableElements[x][y];
                        if (currentElement != $firstCell[0]) {
                            $(currentElement).remove();
                        }
                        _tableElements[x][y] = $firstCell[0];
                        _this.savePosition(_tableElements[x][y], {x, y});
                    }
                }

                _this.renderColGroup($(table));
            }
        }
    },
    deleteTable: function () {
        var table = this._getTable();
        if (table) {
            this.component.remove(table);
        }
    },

    clearTableTheme: function () {
        var table = this._getTable();
        if (table) {
            $R
                .dom(table)
                .removeClass(
                    this.opts.tableClassNames.toString().replace(/[,;]/g, " ")
                );
        }
    },

    clearCellBackground: function () {
        var table = this._getTable();
        if (table) {
            this.selectedCells.removeClass(
                this.opts.cellBackgroundNames.toString().replace(/[,;]/g, " ")
            );
        }
    },

    setTableTheme: function (argus) {
        var table = this._getTable();
        if (table) {
            this.clearTableTheme();
            $R.dom(table).addClass(argus.classValue);
            this.app.api("module.modal.close");
        }
    },

    setCellBackground: function (argus) {
        var table = this._getTable();

        if (table) {
            var cells = this.selectedCells;

            this.clearCellBackground();
            cells.addClass(argus.classValue);

            this.app.api("module.modal.close");
        }
    },

    setCellFontSize: function (argus) {
        var table = this._getTable();
        if (table) {
            var cells = this.selectedCells;
            var size = argus.classValue;
            if (size) {
                cells.css({"font-size": size + "px"});
            } else {
                cells.css({"font-size": ""});
            }
            cells.attr('data-redactor-style-cache', true);

            this.app.api("module.modal.close");
        }
    },

    setTableThemeModal: function () {
        var options = {
            name: "setTableThemeModal",
            title: "Set table theme",
            handle: "save", // optional, command which will be fired on enter pressed
            // optional object
            commands: {
                save: {title: "Save"},
                cancel: {title: "Cancel"}
            }
        };

        this.app.api("module.modal.build", options);
    },

    setCellBackgroundModal: function () {
        var options = {
            name: "setCellBackgroundModal",
            title: "Set cell background",
            handle: "save", // optional, command which will be fired on enter pressed
            // optional object
            commands: {
                save: {title: "Save"},
                cancel: {title: "Cancel"}
            }
        };

        this.app.api("module.modal.build", options);
    },

    setCellFontSizeModal: function () {
        this.cellFontSizes = [12, 16, 18, 34];

        if (typeof this.opts.cellFontSizes != "undefined") {
            this.cellFontSizes = this.opts.cellFontSizes.split(",");
        }
        var options = {
            name: "setCellFontSizeModal",
            title: "Set cell font size",
            handle: "save", // optional, command which will be fired on enter pressed
            // optional object
            commands: {
                save: {title: "Save"},
                cancel: {title: "Cancel"}
            }
        };

        this.app.api("module.modal.build", options);
    },

    // private
    _getTable: function () {
        var current = this.selection.getCurrent();
        var data = this.inspector.parse(current);
        if (data.isTable()) {
            return data.getTable();
        }
    },
    _getComponent: function () {
        var current = this.selection.getCurrent();
        var data = this.inspector.parse(current);
        if (data.isTable()) {
            var table = data.getTable();

            return this.component.create("table", table);
        }
    },
    _observeDropdown: function (dropdown) {
        var table = this._getTable();
        var items = dropdown.getItemsByClass("redactor-table-item-observable");
        var tableItem = dropdown.getItem("insert-table");
        if (table) {
            this._observeItems(items, "enable");
            tableItem.disable();
        } else {
            this._observeItems(items, "disable");
            tableItem.enable();
        }
    },
    _observeItems: function (items, type) {
        for (var i = 0; i < items.length; i++) {
            items[i][type]();
        }
    }
});
