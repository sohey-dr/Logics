use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    pub fn alert(s: &str);
}

#[wasm_bindgen]
pub fn is_expression_valid(word: &str) {
    if is_match_expression_valid(word) {
        let alert_word = word.to_string() + "は不適切な表現です";
        alert(&alert_word);
        return;
    }
}

fn is_match_expression_valid(word: &str) -> bool {
    let invalid_words = vec!["うんち", "うんこ", "おしっこ"];
    for iw in invalid_words {
        if iw.contains(word) {
            return true;
        }
    }

    return false;
}

#[cfg(test)]
mod tests {
    #[test]
    fn is_match_expression_valid() {
        assert_eq!(super::is_match_expression_valid("うんち"), true);
    }

    #[test]
    fn is_not_match_expression_valid() {
        assert_eq!(super::is_match_expression_valid("きれい"), false);
    }
}
