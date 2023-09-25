fn score(word: impl Into<String>) -> usize {
    word.into().replace("a", "").len()
}
fn penalty(word: impl Into<String>) -> usize {
    let word = word.into();
    let base = score(word.clone());
    if word.contains("s") {
        base - 2
    } else {
        base
    }
}
fn high_scoring_words<
    F: Fn(String) -> usize,
    //    R2Fn: Fn(Vec<String>) -> Vec<String>,
    //    R1Fn: Fn(usize) -> R2Fn,
>(
    f: &F,
) -> Box<dyn Fn(usize) -> Box<dyn Fn(Vec<String>) -> Vec<String>>> + '_ {
    Box::new(|higher_than: usize| {
        Box::new(|words: Vec<String>| {
            let mut words = words.clone();
            let w = words
                .into_iter()
                .filter(|word| f(word.clone()) > higher_than)
                .into_iter()
                .collect();
            w
        })
    })
}

fn ranked_words(f: impl Fn(String) -> usize, words: Vec<String>) -> Vec<String> {
    let mut words = words.clone();
    let comp_f = |a: &String, b: &String| {
        let score_a = f(a.clone());
        let score_b = f(b.clone());

        score_a.cmp(&score_b)
    };
    words.sort_by(comp_f);
    words.reverse();
    words
}
fn score_with_bonus(word: impl Into<String>) -> usize {
    let word = word.into();
    let score = score(word.clone());
    if word.contains("c") {
        score + 5
    } else {
        score
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn score_test() {
        let haskell = score("haskell");
        let rust = score("rust");

        assert_eq!(haskell, 6);
        assert_eq!(rust, 4);
    }
    #[test]
    fn ranked_words_test() {
        let words = vec![
            "haskell".to_string(),
            "rust".to_string(),
            "java".to_string(),
            "scala".to_string(),
        ];
        let expected = vec![
            "haskell".to_string(),
            "rust".to_string(),
            "scala".to_string(),
            "java".to_string(),
        ];

        assert_eq!(ranked_words(score, words.clone()), expected);
        assert_eq!(
            ranked_words(score_with_bonus, words.clone()),
            vec![
                "scala".to_string(),
                "haskell".to_string(),
                "rust".to_string(),
                "java".to_string(),
            ]
        );
        assert_eq!(
            ranked_words(penalty, words.clone()),
            vec![
                "haskell".to_string(),
                "java".to_string(),
                "rust".to_string(),
                "scala".to_string(),
            ]
        );
    }
}
