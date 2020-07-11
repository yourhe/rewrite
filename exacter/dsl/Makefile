ANTRL = java -jar ${PWD}/antlr-4.7.1-complete.jar

clear_altrl_tmp:
	rm -rf ${PWD}/.antlr
antrl: clear_altrl_tmp
	$(ANTRL) -Dlanguage=Go -o parser -visitor dsl.g4
build_script: antrl

.PHONY: build_script