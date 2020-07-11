// Generated from /Users/yorhe/Desktop/work/dev/project/2019/wfgz/dr2am-rewrite/exacter/dsl/dsl.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class dslParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, LPAREN=5, RPAREN=6, LBRACK=7, RBRACK=8, 
		LBRACE=9, RBRACE=10, COMMA=11, SEMICOLON=12, ANDAND=13, OROR=14, EQ=15, 
		IF=16, ELSE=17, WHILE=18, BREAK=19, READ=20, WRITE=21, INT=22, REAL=23, 
		Operator_lit=24, Duration_unit=25, VAR=26, PIPE=27, DOT=28, AT=29, Integer=30, 
		RealNumber=31, BooleanLiteral=32, Reference=33, Identifier=34, WS=35, 
		Comment=36, LineComment=37, StringLiteral=38, RegexLiteral=39, DurationLiteral=40, 
		RoundLiteral=41;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_typeDeclaration = 2, RULE_declaration = 3, 
		RULE_expression = 4, RULE_condition = 5, RULE_mapExpr = 6, RULE_mapprimay = 7, 
		RULE_chain = 8, RULE_function = 9, RULE_parameters = 10, RULE_parameter = 11, 
		RULE_primaryExpr = 12, RULE_primary = 13, RULE_lambda = 14, RULE_lambdaparameters = 15;
	public static final String[] ruleNames = {
		"program", "statement", "typeDeclaration", "declaration", "expression", 
		"condition", "mapExpr", "mapprimay", "chain", "function", "parameters", 
		"parameter", "primaryExpr", "primary", "lambda", "lambdaparameters"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'when'", "'map'", "'=>'", "'lambda:'", "'('", "')'", "'['", "']'", 
		"'{'", "'}'", "','", "';'", "'&&'", "'||'", "'='", "'if'", "'else'", "'while'", 
		"'break'", "'read'", "'write'", "'int'", "'real'", null, null, "'var'", 
		"'|'", "'.'", "'@'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, "LPAREN", "RPAREN", "LBRACK", "RBRACK", 
		"LBRACE", "RBRACE", "COMMA", "SEMICOLON", "ANDAND", "OROR", "EQ", "IF", 
		"ELSE", "WHILE", "BREAK", "READ", "WRITE", "INT", "REAL", "Operator_lit", 
		"Duration_unit", "VAR", "PIPE", "DOT", "AT", "Integer", "RealNumber", 
		"BooleanLiteral", "Reference", "Identifier", "WS", "Comment", "LineComment", 
		"StringLiteral", "RegexLiteral", "DurationLiteral", "RoundLiteral"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "dsl.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public dslParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public TerminalNode EOF() { return getToken(dslParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			statement();
			setState(36);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << LineComment))) != 0)) {
				{
				{
				setState(33);
				statement();
				}
				}
				setState(38);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(39);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeDeclarationContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(dslParser.VAR, 0); }
		public List<TerminalNode> Identifier() { return getTokens(dslParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(dslParser.Identifier, i);
		}
		public TypeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDeclaration; }
	}

	public final TypeDeclarationContext typeDeclaration() throws RecognitionException {
		TypeDeclarationContext _localctx = new TypeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(VAR);
			setState(44);
			match(Identifier);
			setState(45);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(dslParser.VAR, 0); }
		public TerminalNode Identifier() { return getToken(dslParser.Identifier, 0); }
		public TerminalNode EQ() { return getToken(dslParser.EQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(VAR);
			setState(48);
			match(Identifier);
			setState(49);
			match(EQ);
			setState(50);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public ConditionContext condition() {
			return getRuleContext(ConditionContext.class,0);
		}
		public MapExprContext mapExpr() {
			return getRuleContext(MapExprContext.class,0);
		}
		public TerminalNode LineComment() { return getToken(dslParser.LineComment, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_expression);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				condition();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
				mapExpr();
				}
				break;
			case LineComment:
				enterOuterAlt(_localctx, 3);
				{
				setState(54);
				match(LineComment);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionContext extends ParserRuleContext {
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_condition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(T__0);
			setState(58);
			primaryExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MapExprContext extends ParserRuleContext {
		public List<MapprimayContext> mapprimay() {
			return getRuleContexts(MapprimayContext.class);
		}
		public MapprimayContext mapprimay(int i) {
			return getRuleContext(MapprimayContext.class,i);
		}
		public MapExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapExpr; }
	}

	public final MapExprContext mapExpr() throws RecognitionException {
		MapExprContext _localctx = new MapExprContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_mapExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(T__1);
			setState(62); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(61);
				mapprimay();
				}
				}
				setState(64); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Reference );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MapprimayContext extends ParserRuleContext {
		public TerminalNode Reference() { return getToken(dslParser.Reference, 0); }
		public MapprimayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapprimay; }
	}

	public final MapprimayContext mapprimay() throws RecognitionException {
		MapprimayContext _localctx = new MapprimayContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_mapprimay);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(Reference);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ChainContext extends ParserRuleContext {
		public TerminalNode AT() { return getToken(dslParser.AT, 0); }
		public FunctionContext function() {
			return getRuleContext(FunctionContext.class,0);
		}
		public TerminalNode PIPE() { return getToken(dslParser.PIPE, 0); }
		public List<ChainContext> chain() {
			return getRuleContexts(ChainContext.class);
		}
		public ChainContext chain(int i) {
			return getRuleContext(ChainContext.class,i);
		}
		public TerminalNode DOT() { return getToken(dslParser.DOT, 0); }
		public TerminalNode Identifier() { return getToken(dslParser.Identifier, 0); }
		public ChainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chain; }
	}

	public final ChainContext chain() throws RecognitionException {
		ChainContext _localctx = new ChainContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_chain);
		try {
			int _alt;
			setState(96);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				match(AT);
				setState(69);
				function();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(70);
				match(PIPE);
				setState(71);
				function();
				setState(75);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(72);
						chain();
						}
						} 
					}
					setState(77);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(78);
				match(DOT);
				setState(79);
				function();
				setState(83);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(80);
						chain();
						}
						} 
					}
					setState(85);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(86);
				match(DOT);
				setState(87);
				match(Identifier);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(88);
				match(DOT);
				setState(89);
				match(Identifier);
				setState(93);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(90);
						chain();
						}
						} 
					}
					setState(95);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(dslParser.Identifier, 0); }
		public TerminalNode LPAREN() { return getToken(dslParser.LPAREN, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(dslParser.RPAREN, 0); }
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_function);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(Identifier);
			setState(99);
			match(LPAREN);
			setState(100);
			parameters();
			setState(101);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametersContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_parameters);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(103);
					parameter();
					setState(104);
					match(COMMA);
					}
					} 
				}
				setState(110);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << LineComment))) != 0)) {
				{
				setState(111);
				parameter();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimaryExprContext extends ParserRuleContext {
		public List<PrimaryContext> primary() {
			return getRuleContexts(PrimaryContext.class);
		}
		public PrimaryContext primary(int i) {
			return getRuleContext(PrimaryContext.class,i);
		}
		public List<TerminalNode> Operator_lit() { return getTokens(dslParser.Operator_lit); }
		public TerminalNode Operator_lit(int i) {
			return getToken(dslParser.Operator_lit, i);
		}
		public PrimaryExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpr; }
	}

	public final PrimaryExprContext primaryExpr() throws RecognitionException {
		PrimaryExprContext _localctx = new PrimaryExprContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_primaryExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			primary();
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Operator_lit) {
				{
				{
				setState(117);
				match(Operator_lit);
				setState(118);
				primary();
				}
				}
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimaryContext extends ParserRuleContext {
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public TerminalNode RealNumber() { return getToken(dslParser.RealNumber, 0); }
		public TerminalNode Operator_lit() { return getToken(dslParser.Operator_lit, 0); }
		public TerminalNode Integer() { return getToken(dslParser.Integer, 0); }
		public TerminalNode StringLiteral() { return getToken(dslParser.StringLiteral, 0); }
		public TerminalNode Reference() { return getToken(dslParser.Reference, 0); }
		public TerminalNode DurationLiteral() { return getToken(dslParser.DurationLiteral, 0); }
		public TerminalNode BooleanLiteral() { return getToken(dslParser.BooleanLiteral, 0); }
		public TerminalNode Identifier() { return getToken(dslParser.Identifier, 0); }
		public TerminalNode RegexLiteral() { return getToken(dslParser.RegexLiteral, 0); }
		public TerminalNode RoundLiteral() { return getToken(dslParser.RoundLiteral, 0); }
		public PrimaryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary; }
	}

	public final PrimaryContext primary() throws RecognitionException {
		PrimaryContext _localctx = new PrimaryContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_primary);
		try {
			setState(138);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(LPAREN);
				setState(125);
				primaryExpr();
				setState(126);
				match(RPAREN);
				}
				break;
			case RealNumber:
				enterOuterAlt(_localctx, 2);
				{
				setState(128);
				match(RealNumber);
				}
				break;
			case Operator_lit:
				enterOuterAlt(_localctx, 3);
				{
				setState(129);
				match(Operator_lit);
				}
				break;
			case Integer:
				enterOuterAlt(_localctx, 4);
				{
				setState(130);
				match(Integer);
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 5);
				{
				setState(131);
				match(StringLiteral);
				}
				break;
			case Reference:
				enterOuterAlt(_localctx, 6);
				{
				setState(132);
				match(Reference);
				}
				break;
			case DurationLiteral:
				enterOuterAlt(_localctx, 7);
				{
				setState(133);
				match(DurationLiteral);
				}
				break;
			case BooleanLiteral:
				enterOuterAlt(_localctx, 8);
				{
				setState(134);
				match(BooleanLiteral);
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 9);
				{
				setState(135);
				match(Identifier);
				}
				break;
			case RegexLiteral:
				enterOuterAlt(_localctx, 10);
				{
				setState(136);
				match(RegexLiteral);
				}
				break;
			case RoundLiteral:
				enterOuterAlt(_localctx, 11);
				{
				setState(137);
				match(RoundLiteral);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LambdaContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(dslParser.LPAREN, 0); }
		public LambdaparametersContext lambdaparameters() {
			return getRuleContext(LambdaparametersContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(dslParser.RPAREN, 0); }
		public TerminalNode LBRACE() { return getToken(dslParser.LBRACE, 0); }
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(dslParser.RBRACE, 0); }
		public LambdaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lambda; }
	}

	public final LambdaContext lambda() throws RecognitionException {
		LambdaContext _localctx = new LambdaContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_lambda);
		try {
			setState(150);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(140);
				match(LPAREN);
				setState(141);
				lambdaparameters();
				setState(142);
				match(RPAREN);
				setState(143);
				match(T__2);
				setState(144);
				match(LBRACE);
				setState(145);
				primaryExpr();
				setState(146);
				match(RBRACE);
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				match(T__3);
				setState(149);
				primaryExpr();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LambdaparametersContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(dslParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(dslParser.Identifier, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public LambdaparametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lambdaparameters; }
	}

	public final LambdaparametersContext lambdaparameters() throws RecognitionException {
		LambdaparametersContext _localctx = new LambdaparametersContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_lambdaparameters);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(154);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case Identifier:
						{
						setState(152);
						match(Identifier);
						}
						break;
					case T__0:
					case T__1:
					case LineComment:
						{
						setState(153);
						expression();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(156);
					match(COMMA);
					}
					} 
				}
				setState(161);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			}
			setState(164);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				{
				setState(162);
				match(Identifier);
				}
				break;
			case T__0:
			case T__1:
			case LineComment:
				{
				setState(163);
				expression();
				}
				break;
			case RPAREN:
				break;
			default:
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3+\u00a9\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\7"+
		"\2%\n\2\f\2\16\2(\13\2\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5"+
		"\3\5\3\6\3\6\3\6\5\6:\n\6\3\7\3\7\3\7\3\b\3\b\6\bA\n\b\r\b\16\bB\3\t\3"+
		"\t\3\n\3\n\3\n\3\n\3\n\7\nL\n\n\f\n\16\nO\13\n\3\n\3\n\3\n\7\nT\n\n\f"+
		"\n\16\nW\13\n\3\n\3\n\3\n\3\n\3\n\7\n^\n\n\f\n\16\na\13\n\5\nc\n\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\f\3\f\3\f\7\fm\n\f\f\f\16\fp\13\f\3\f\5\fs\n\f"+
		"\3\r\3\r\3\16\3\16\3\16\7\16z\n\16\f\16\16\16}\13\16\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u008d\n\17"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0099\n\20\3\21"+
		"\3\21\5\21\u009d\n\21\3\21\7\21\u00a0\n\21\f\21\16\21\u00a3\13\21\3\21"+
		"\3\21\5\21\u00a7\n\21\3\21\2\2\22\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \2\2\2\u00b5\2\"\3\2\2\2\4+\3\2\2\2\6-\3\2\2\2\b\61\3\2\2\2\n9\3\2"+
		"\2\2\f;\3\2\2\2\16>\3\2\2\2\20D\3\2\2\2\22b\3\2\2\2\24d\3\2\2\2\26n\3"+
		"\2\2\2\30t\3\2\2\2\32v\3\2\2\2\34\u008c\3\2\2\2\36\u0098\3\2\2\2 \u00a1"+
		"\3\2\2\2\"&\5\4\3\2#%\5\4\3\2$#\3\2\2\2%(\3\2\2\2&$\3\2\2\2&\'\3\2\2\2"+
		"\')\3\2\2\2(&\3\2\2\2)*\7\2\2\3*\3\3\2\2\2+,\5\n\6\2,\5\3\2\2\2-.\7\34"+
		"\2\2./\7$\2\2/\60\7$\2\2\60\7\3\2\2\2\61\62\7\34\2\2\62\63\7$\2\2\63\64"+
		"\7\21\2\2\64\65\5\n\6\2\65\t\3\2\2\2\66:\5\f\7\2\67:\5\16\b\28:\7\'\2"+
		"\29\66\3\2\2\29\67\3\2\2\298\3\2\2\2:\13\3\2\2\2;<\7\3\2\2<=\5\32\16\2"+
		"=\r\3\2\2\2>@\7\4\2\2?A\5\20\t\2@?\3\2\2\2AB\3\2\2\2B@\3\2\2\2BC\3\2\2"+
		"\2C\17\3\2\2\2DE\7#\2\2E\21\3\2\2\2FG\7\37\2\2Gc\5\24\13\2HI\7\35\2\2"+
		"IM\5\24\13\2JL\5\22\n\2KJ\3\2\2\2LO\3\2\2\2MK\3\2\2\2MN\3\2\2\2Nc\3\2"+
		"\2\2OM\3\2\2\2PQ\7\36\2\2QU\5\24\13\2RT\5\22\n\2SR\3\2\2\2TW\3\2\2\2U"+
		"S\3\2\2\2UV\3\2\2\2Vc\3\2\2\2WU\3\2\2\2XY\7\36\2\2Yc\7$\2\2Z[\7\36\2\2"+
		"[_\7$\2\2\\^\5\22\n\2]\\\3\2\2\2^a\3\2\2\2_]\3\2\2\2_`\3\2\2\2`c\3\2\2"+
		"\2a_\3\2\2\2bF\3\2\2\2bH\3\2\2\2bP\3\2\2\2bX\3\2\2\2bZ\3\2\2\2c\23\3\2"+
		"\2\2de\7$\2\2ef\7\7\2\2fg\5\26\f\2gh\7\b\2\2h\25\3\2\2\2ij\5\30\r\2jk"+
		"\7\r\2\2km\3\2\2\2li\3\2\2\2mp\3\2\2\2nl\3\2\2\2no\3\2\2\2or\3\2\2\2p"+
		"n\3\2\2\2qs\5\30\r\2rq\3\2\2\2rs\3\2\2\2s\27\3\2\2\2tu\5\n\6\2u\31\3\2"+
		"\2\2v{\5\34\17\2wx\7\32\2\2xz\5\34\17\2yw\3\2\2\2z}\3\2\2\2{y\3\2\2\2"+
		"{|\3\2\2\2|\33\3\2\2\2}{\3\2\2\2~\177\7\7\2\2\177\u0080\5\32\16\2\u0080"+
		"\u0081\7\b\2\2\u0081\u008d\3\2\2\2\u0082\u008d\7!\2\2\u0083\u008d\7\32"+
		"\2\2\u0084\u008d\7 \2\2\u0085\u008d\7(\2\2\u0086\u008d\7#\2\2\u0087\u008d"+
		"\7*\2\2\u0088\u008d\7\"\2\2\u0089\u008d\7$\2\2\u008a\u008d\7)\2\2\u008b"+
		"\u008d\7+\2\2\u008c~\3\2\2\2\u008c\u0082\3\2\2\2\u008c\u0083\3\2\2\2\u008c"+
		"\u0084\3\2\2\2\u008c\u0085\3\2\2\2\u008c\u0086\3\2\2\2\u008c\u0087\3\2"+
		"\2\2\u008c\u0088\3\2\2\2\u008c\u0089\3\2\2\2\u008c\u008a\3\2\2\2\u008c"+
		"\u008b\3\2\2\2\u008d\35\3\2\2\2\u008e\u008f\7\7\2\2\u008f\u0090\5 \21"+
		"\2\u0090\u0091\7\b\2\2\u0091\u0092\7\5\2\2\u0092\u0093\7\13\2\2\u0093"+
		"\u0094\5\32\16\2\u0094\u0095\7\f\2\2\u0095\u0099\3\2\2\2\u0096\u0097\7"+
		"\6\2\2\u0097\u0099\5\32\16\2\u0098\u008e\3\2\2\2\u0098\u0096\3\2\2\2\u0099"+
		"\37\3\2\2\2\u009a\u009d\7$\2\2\u009b\u009d\5\n\6\2\u009c\u009a\3\2\2\2"+
		"\u009c\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u00a0\7\r\2\2\u009f\u009c"+
		"\3\2\2\2\u00a0\u00a3\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2"+
		"\u00a6\3\2\2\2\u00a3\u00a1\3\2\2\2\u00a4\u00a7\7$\2\2\u00a5\u00a7\5\n"+
		"\6\2\u00a6\u00a4\3\2\2\2\u00a6\u00a5\3\2\2\2\u00a6\u00a7\3\2\2\2\u00a7"+
		"!\3\2\2\2\21&9BMU_bnr{\u008c\u0098\u009c\u00a1\u00a6";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}