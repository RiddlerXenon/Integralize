import { parse } from 'mathjs';

export const toLatex = (expression) => {
    if (!expression) {
        return 'x'; // Возвращаем 'x' если выражение пустое
    }
    try {
        const node = parse(expression);
        return node.toTex({ parenthesis: 'keep', implicit: 'show' });
    } catch (error) {
        console.error('Ошибка при преобразовании в LaTeX:', error);
        return expression;
    }
};
