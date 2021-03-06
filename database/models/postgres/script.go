// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
)

// Script is an object representing the database table.
type Script struct {
	ID             string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	ScriptID       string     `boil:"script_id" json:"script_id" toml:"script_id" yaml:"script_id"`
	ScriptName     string     `boil:"script_name" json:"script_name" toml:"script_name" yaml:"script_name"`
	ScriptPath     string     `boil:"script_path" json:"script_path" toml:"script_path" yaml:"script_path"`
	ScriptData     null.Bytes `boil:"script_data" json:"script_data,omitempty" toml:"script_data" yaml:"script_data,omitempty"`
	LastExecutedAt null.Time  `boil:"last_executed_at" json:"last_executed_at,omitempty" toml:"last_executed_at" yaml:"last_executed_at,omitempty"`
	CreatedAt      null.Time  `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *scriptR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scriptL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ScriptColumns = struct {
	ID             string
	ScriptID       string
	ScriptName     string
	ScriptPath     string
	ScriptData     string
	LastExecutedAt string
	CreatedAt      string
}{
	ID:             "id",
	ScriptID:       "script_id",
	ScriptName:     "script_name",
	ScriptPath:     "script_path",
	ScriptData:     "script_data",
	LastExecutedAt: "last_executed_at",
	CreatedAt:      "created_at",
}

// Generated where

type whereHelpernull_Bytes struct{ field string }

func (w whereHelpernull_Bytes) EQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bytes) NEQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bytes) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bytes) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Bytes) LT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bytes) LTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bytes) GT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bytes) GTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ScriptWhere = struct {
	ID             whereHelperstring
	ScriptID       whereHelperstring
	ScriptName     whereHelperstring
	ScriptPath     whereHelperstring
	ScriptData     whereHelpernull_Bytes
	LastExecutedAt whereHelpernull_Time
	CreatedAt      whereHelpernull_Time
}{
	ID:             whereHelperstring{field: "\"script\".\"id\""},
	ScriptID:       whereHelperstring{field: "\"script\".\"script_id\""},
	ScriptName:     whereHelperstring{field: "\"script\".\"script_name\""},
	ScriptPath:     whereHelperstring{field: "\"script\".\"script_path\""},
	ScriptData:     whereHelpernull_Bytes{field: "\"script\".\"script_data\""},
	LastExecutedAt: whereHelpernull_Time{field: "\"script\".\"last_executed_at\""},
	CreatedAt:      whereHelpernull_Time{field: "\"script\".\"created_at\""},
}

// ScriptRels is where relationship names are stored.
var ScriptRels = struct {
	ScriptExecutions string
}{
	ScriptExecutions: "ScriptExecutions",
}

// scriptR is where relationships are stored.
type scriptR struct {
	ScriptExecutions ScriptExecutionSlice
}

// NewStruct creates a new relationship struct
func (*scriptR) NewStruct() *scriptR {
	return &scriptR{}
}

// scriptL is where Load methods for each relationship are stored.
type scriptL struct{}

var (
	scriptAllColumns            = []string{"id", "script_id", "script_name", "script_path", "script_data", "last_executed_at", "created_at"}
	scriptColumnsWithoutDefault = []string{"script_id", "script_name", "script_path", "script_data"}
	scriptColumnsWithDefault    = []string{"id", "last_executed_at", "created_at"}
	scriptPrimaryKeyColumns     = []string{"id"}
)

type (
	// ScriptSlice is an alias for a slice of pointers to Script.
	// This should generally be used opposed to []Script.
	ScriptSlice []*Script
	// ScriptHook is the signature for custom Script hook methods
	ScriptHook func(context.Context, boil.ContextExecutor, *Script) error

	scriptQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	scriptType                 = reflect.TypeOf(&Script{})
	scriptMapping              = queries.MakeStructMapping(scriptType)
	scriptPrimaryKeyMapping, _ = queries.BindMapping(scriptType, scriptMapping, scriptPrimaryKeyColumns)
	scriptInsertCacheMut       sync.RWMutex
	scriptInsertCache          = make(map[string]insertCache)
	scriptUpdateCacheMut       sync.RWMutex
	scriptUpdateCache          = make(map[string]updateCache)
	scriptUpsertCacheMut       sync.RWMutex
	scriptUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var scriptBeforeInsertHooks []ScriptHook
var scriptBeforeUpdateHooks []ScriptHook
var scriptBeforeDeleteHooks []ScriptHook
var scriptBeforeUpsertHooks []ScriptHook

var scriptAfterInsertHooks []ScriptHook
var scriptAfterSelectHooks []ScriptHook
var scriptAfterUpdateHooks []ScriptHook
var scriptAfterDeleteHooks []ScriptHook
var scriptAfterUpsertHooks []ScriptHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Script) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Script) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Script) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Script) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Script) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Script) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Script) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Script) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Script) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddScriptHook registers your hook function for all future operations.
func AddScriptHook(hookPoint boil.HookPoint, scriptHook ScriptHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		scriptBeforeInsertHooks = append(scriptBeforeInsertHooks, scriptHook)
	case boil.BeforeUpdateHook:
		scriptBeforeUpdateHooks = append(scriptBeforeUpdateHooks, scriptHook)
	case boil.BeforeDeleteHook:
		scriptBeforeDeleteHooks = append(scriptBeforeDeleteHooks, scriptHook)
	case boil.BeforeUpsertHook:
		scriptBeforeUpsertHooks = append(scriptBeforeUpsertHooks, scriptHook)
	case boil.AfterInsertHook:
		scriptAfterInsertHooks = append(scriptAfterInsertHooks, scriptHook)
	case boil.AfterSelectHook:
		scriptAfterSelectHooks = append(scriptAfterSelectHooks, scriptHook)
	case boil.AfterUpdateHook:
		scriptAfterUpdateHooks = append(scriptAfterUpdateHooks, scriptHook)
	case boil.AfterDeleteHook:
		scriptAfterDeleteHooks = append(scriptAfterDeleteHooks, scriptHook)
	case boil.AfterUpsertHook:
		scriptAfterUpsertHooks = append(scriptAfterUpsertHooks, scriptHook)
	}
}

// One returns a single script record from the query.
func (q scriptQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Script, error) {
	o := &Script{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: failed to execute a one query for script")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Script records from the query.
func (q scriptQuery) All(ctx context.Context, exec boil.ContextExecutor) (ScriptSlice, error) {
	var o []*Script

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: failed to assign all query results to Script slice")
	}

	if len(scriptAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Script records in the query.
func (q scriptQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to count script rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q scriptQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgres: failed to check if script exists")
	}

	return count > 0, nil
}

// ScriptExecutions retrieves all the script_execution's ScriptExecutions with an executor.
func (o *Script) ScriptExecutions(mods ...qm.QueryMod) scriptExecutionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"script_execution\".\"script_id\"=?", o.ID),
	)

	query := ScriptExecutions(queryMods...)
	queries.SetFrom(query.Query, "\"script_execution\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"script_execution\".*"})
	}

	return query
}

// LoadScriptExecutions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (scriptL) LoadScriptExecutions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeScript interface{}, mods queries.Applicator) error {
	var slice []*Script
	var object *Script

	if singular {
		object = maybeScript.(*Script)
	} else {
		slice = *maybeScript.(*[]*Script)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &scriptR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &scriptR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`script_execution`), qm.WhereIn(`script_execution.script_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load script_execution")
	}

	var resultSlice []*ScriptExecution
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice script_execution")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on script_execution")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for script_execution")
	}

	if len(scriptExecutionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ScriptExecutions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &scriptExecutionR{}
			}
			foreign.R.Script = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ScriptID) {
				local.R.ScriptExecutions = append(local.R.ScriptExecutions, foreign)
				if foreign.R == nil {
					foreign.R = &scriptExecutionR{}
				}
				foreign.R.Script = local
				break
			}
		}
	}

	return nil
}

// AddScriptExecutions adds the given related objects to the existing relationships
// of the script, optionally inserting them as new records.
// Appends related to o.R.ScriptExecutions.
// Sets related.R.Script appropriately.
func (o *Script) AddScriptExecutions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ScriptExecution) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ScriptID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"script_execution\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"script_id"}),
				strmangle.WhereClause("\"", "\"", 2, scriptExecutionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ScriptID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &scriptR{
			ScriptExecutions: related,
		}
	} else {
		o.R.ScriptExecutions = append(o.R.ScriptExecutions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &scriptExecutionR{
				Script: o,
			}
		} else {
			rel.R.Script = o
		}
	}
	return nil
}

// SetScriptExecutions removes all previously related items of the
// script replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Script's ScriptExecutions accordingly.
// Replaces o.R.ScriptExecutions with related.
// Sets related.R.Script's ScriptExecutions accordingly.
func (o *Script) SetScriptExecutions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ScriptExecution) error {
	query := "update \"script_execution\" set \"script_id\" = null where \"script_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ScriptExecutions {
			queries.SetScanner(&rel.ScriptID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Script = nil
		}

		o.R.ScriptExecutions = nil
	}
	return o.AddScriptExecutions(ctx, exec, insert, related...)
}

// RemoveScriptExecutions relationships from objects passed in.
// Removes related items from R.ScriptExecutions (uses pointer comparison, removal does not keep order)
// Sets related.R.Script.
func (o *Script) RemoveScriptExecutions(ctx context.Context, exec boil.ContextExecutor, related ...*ScriptExecution) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ScriptID, nil)
		if rel.R != nil {
			rel.R.Script = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("script_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ScriptExecutions {
			if rel != ri {
				continue
			}

			ln := len(o.R.ScriptExecutions)
			if ln > 1 && i < ln-1 {
				o.R.ScriptExecutions[i] = o.R.ScriptExecutions[ln-1]
			}
			o.R.ScriptExecutions = o.R.ScriptExecutions[:ln-1]
			break
		}
	}

	return nil
}

// Scripts retrieves all the records using an executor.
func Scripts(mods ...qm.QueryMod) scriptQuery {
	mods = append(mods, qm.From("\"script\""))
	return scriptQuery{NewQuery(mods...)}
}

// FindScript retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindScript(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Script, error) {
	scriptObj := &Script{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"script\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, scriptObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: unable to select from script")
	}

	return scriptObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Script) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no script provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(scriptColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	scriptInsertCacheMut.RLock()
	cache, cached := scriptInsertCache[key]
	scriptInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			scriptAllColumns,
			scriptColumnsWithDefault,
			scriptColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(scriptType, scriptMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(scriptType, scriptMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"script\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"script\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "postgres: unable to insert into script")
	}

	if !cached {
		scriptInsertCacheMut.Lock()
		scriptInsertCache[key] = cache
		scriptInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Script.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Script) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	scriptUpdateCacheMut.RLock()
	cache, cached := scriptUpdateCache[key]
	scriptUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			scriptAllColumns,
			scriptPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("postgres: unable to update script, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"script\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, scriptPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(scriptType, scriptMapping, append(wl, scriptPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update script row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by update for script")
	}

	if !cached {
		scriptUpdateCacheMut.Lock()
		scriptUpdateCache[key] = cache
		scriptUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q scriptQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all for script")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected for script")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ScriptSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("postgres: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"script\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, scriptPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all in script slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected all in update all script")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Script) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no script provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(scriptColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	scriptUpsertCacheMut.RLock()
	cache, cached := scriptUpsertCache[key]
	scriptUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			scriptAllColumns,
			scriptColumnsWithDefault,
			scriptColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			scriptAllColumns,
			scriptPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgres: unable to upsert script, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(scriptPrimaryKeyColumns))
			copy(conflict, scriptPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"script\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(scriptType, scriptMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(scriptType, scriptMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "postgres: unable to upsert script")
	}

	if !cached {
		scriptUpsertCacheMut.Lock()
		scriptUpsertCache[key] = cache
		scriptUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Script record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Script) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgres: no Script provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), scriptPrimaryKeyMapping)
	sql := "DELETE FROM \"script\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete from script")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by delete for script")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q scriptQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgres: no scriptQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from script")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for script")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ScriptSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(scriptBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"script\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scriptPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from script slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for script")
	}

	if len(scriptAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Script) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindScript(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScriptSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ScriptSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"script\".* FROM \"script\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scriptPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgres: unable to reload all in ScriptSlice")
	}

	*o = slice

	return nil
}

// ScriptExists checks if the Script row exists.
func ScriptExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"script\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgres: unable to check if script exists")
	}

	return exists, nil
}
