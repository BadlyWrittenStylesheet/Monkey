package evaluator

import (

    "fmt"
	"BadlyWrittenStylesheet/Monkey/monkey/ast"
	"BadlyWrittenStylesheet/Monkey/monkey/object"
)

var (
	TRUE	= &object.Boolean{Value: true}
	FALSE	= &object.Boolean{Value: false}
	NULL	= &object.Null{}
)

// func Eval(node ast.Node, env *object.Enviroment) object.Object {
//     fmt.Println("Evaluating node:", node)
//     fmt.Println("Current environment:", env)

//     switch node := node.(type) {
//     case *ast.Program:
//         return evalProgram(node, env)
//     case *ast.ExpressionStatement:
//         return Eval(node.Expression, env)
//     case *ast.PrefixExpression:
//         right := Eval(node.Right, env)
//         if isError(right) {
//             return right
//         }
//         return evalPrefixExpression(node.Operator, right)
//     case *ast.InfixExpression:
//         left := Eval(node.Left, env)
//         if isError(left) {
//             return left
//         }
//         right := Eval(node.Right, env)
//         if isError(right) {
//             return right
//         }
//         return evalInfixExpression(node.Operator, left, right)
//     case *ast.BlockStatement:
//         return evalBlockStatement(node, env)
//     case *ast.IfExpression:
//         return evalIfExpression(node, env)
//     case *ast.ReturnStatement:
//         val := Eval(node.ReturnValue, env)
//         if isError(val) {
//             return val
//         }
//         return &object.ReturnValue{Value: val}
//     case *ast.FunctionLiteral:
//         params := node.Parameters
//         body := node.Body
//         return &object.Function{Parameters: params, Env: env, Body: body}
//     case *ast.CallExpression:
//         function := Eval(node.Function, env)
//         if isError(function) {
//             return function
//         }
//         args := evalExpressions(node.Arguments, env)
//         if len(args) == 1 && isError(args[0]) {
//             return args[0]
//         }
//         result := applyFunction(function, args)
//         fmt.Println("Result after function application:", result)
//         return result
//     case *ast.LetStatement:
//         val := Eval(node.Value, env)
//         if isError(val) {
//             return val
//         }
//         env.Set(node.Name.Value, val)
//     case *ast.Identifier:
//         return evalIdentifier(node, env)
//     case *ast.IntegerLiteral:
//         return &object.Integer{Value: node.Value}
//     case *ast.Boolean:
//         return nativeBoolToBooleanObject(node.Value)
//     }

//     return nil
// }


func Eval(node ast.Node, env *object.Enviroment) object.Object {
	switch node := node.(type) {

	// Stmts
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
        if isError(left) {
            return left
        }
		right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
    case *ast.ReturnStatement:
        val := Eval(node.ReturnValue, env)
        if isError(val) {
            return val
        }
        return &object.ReturnValue{Value: val}
    case *ast.FunctionLiteral:
        params := node.Parameters
        body := node.Body
        return &object.Function{Parameters: params, Env: env, Body: body}
    case *ast.CallExpression:
        fmt.Println("!1", node)
        fmt.Println("!2", env)
        function := Eval(node.Function, env)
        fmt.Println(function)
        if isError(function) {
            return function
        }
        args := evalExpressions(node.Arguments, env)
        if len(args) == 1 && isError(args[0]) {
            return args[0]
        }
        return applyFunction(function, args)
    case *ast.LetStatement:
        val := Eval(node.Value, env)
        if isError(val) {
            return val
        }
        env.Set(node.Name.Value, val)

    case *ast.Identifier:
        return evalIdentifier(node, env)
	// Exprss
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}

// func applyFunction(fn object.Object, args []object.Object) object.Object {
//     function, ok := fn.(*object.Function)
//     if !ok {
//         return newError("not a function: %s", fn.Type())
//     }

    // extendedEnv := extendFunctionEnv(function, args)
    // evaluated := Eval(function.Body, extendedEnv)
    // return unwrapReturnValue(evaluated)
// }

func applyFunction(fn object.Object, args []object.Object) object.Object {
    fmt.Println("Applying function:", fn)
    function, ok := fn.(*object.Function)
    if !ok {
        return newError("not a function: %s", fn.Type())
    }

    extendedEnv := extendFunctionEnv(function, args)
    fmt.Println("Extended environment for function:", extendedEnv)
    evaluated := Eval(function.Body, extendedEnv)
    fmt.Println("Evaluated function body:", evaluated)
    return unwrapReturnValue(evaluated)
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Enviroment {
    env := object.NewEnclosedEnviroment(fn.Env)
    for paramIdx, param := range fn.Parameters {
        env.Set(param.Value, args[paramIdx])
    }
    return env
}

func unwrapReturnValue(obj object.Object) object.Object {
    if returnValue, ok := obj.(*object.ReturnValue); ok {
        return returnValue.Value
    }

    return obj
}

func evalExpressions(exps []ast.Expression, env *object.Enviroment) []object.Object {
    var result []object.Object

    for _, e := range exps {
        evaluated := Eval(e, env)
        if isError(evaluated) {
            return []object.Object{evaluated}
        }
        result = append(result, evaluated)
    }

    return result
}

func isError(obj object.Object) bool {
    if obj != nil {
        return obj.Type() == object.ERROR_OBJ
    }
    return false
}

func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Enviroment) object.Object {
    condition := Eval(ie.Condition, env)
    if isError(condition) {
        return condition
    }

    fmt.Println("Condition evaluated to:", condition) // Debug print

    if isTruthy(condition) {
        result := Eval(ie.Consequence, env)
        fmt.Println("Consequence:", ie.Consequence, ". Evaluated:", Eval(ie.Consequence, env))
        fmt.Println("Consequence evaluated to:", ie.Consequence, result) // Debug print
        return result
    } else if ie.Alternative != nil {
        result := Eval(ie.Alternative, env)
        fmt.Println("Alternative evaluated to:", result) // Debug print
        return result
    } else {
        fmt.Println("No alternative, returning NULL") // Debug print
        return NULL
    }
}


// func evalIfExpression(ie *ast.IfExpression, env *object.Enviroment) object.Object {
// 	condition := Eval(ie.Condition, env)
//     if isError(condition) {
//         return condition
//     }

// 	if isTruthy(condition) {
// 		return Eval(ie.Consequence, env)
// 	} else if ie.Alternative != nil {
// 		return Eval(ie.Alternative, env)
// 	} else {
// 		return NULL
// 	}
// }

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalIdentifier(node *ast.Identifier, env *object.Enviroment) object.Object {
    val, ok := env.Get(node.Value)
    if !ok {
        return newError("identifier not found: " + node.Value)
    }

    return val
}


func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
    case left.Type() != right.Type():
        return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}


func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
        return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
        return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmts []ast.Statement, env *object.Enviroment) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement, env)

        if returnValue, ok := result.(*object.ReturnValue); ok {
            return returnValue.Value
        }
	}

	return result
}

func evalProgram(program *ast.Program, env *object.Enviroment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

        switch result := result.(type) {
        case *object.ReturnValue:
            return result.Value
        case *object.Error:
            return result
        }
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Enviroment) object.Object {
    var result object.Object

    for _, statement := range block.Statements {
        result = Eval(statement, env)

        if result != nil {
            rt := result.Type()
            if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
                return result
            }
        }
    }

    return result
}

